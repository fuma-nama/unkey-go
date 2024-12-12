import { remark } from "remark";
import remarkStringify from "remark-stringify";
import { visit } from "unist-util-visit";
import { toString as mdastToString } from "mdast-util-to-string";
import matter = require("gray-matter");
import FastGlob = require("fast-glob");
import { readFile, writeFile } from "fs/promises";
import * as path from "node:path";
import { Transformer } from "unified";
import { Root } from "mdast";

const processor = remark().use(remarkMap).use(remarkStringify)

async function main() {
    const files = await FastGlob('../sdks/**/*.md')

    const promises = files.map(async file => {
        const parsed = path.parse(file)
        if (parsed.name.endsWith('.out')) return;

        const content = await readFile(file)
        const out = await mapPage(content.toString())

        // TODO: support to replace original files directly via command arguments
        const outPath = path.join(parsed.dir, `${parsed.name}.out${parsed.ext}`)

        await writeFile(outPath, out)
        console.log('output:', outPath)
    })

    await Promise.all(promises)
}

function remarkMap(): Transformer<Root> {
    return async (tree, file) => {
        let title = 'Untitled'

        visit(tree, ['heading', 'paragraph'], node => {
            if (node.type === 'heading' && node.depth === 1) {
                title = mdastToString(node).trim()
                return 'skip'
            }

            if (title && node.type === 'paragraph') {
                const str = `(${title})`
                if (mdastToString(node).trim() === str) {

                    node.children = []
                    return 'skip'
                }
            }
        })

        file.data.title = title
    }
}

async function mapPage(page: string) {
    const result = await processor.process(page)

    return matter.stringify(String(result.value), {
        title: result.data.title,
    })
}

void main()