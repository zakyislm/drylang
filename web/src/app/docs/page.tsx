import fs from 'fs'
import path from 'path'
import { notFound } from 'next/navigation'
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import { DOCS_STRUCTURE } from './config'
import Link from 'next/link'
import { ArrowLeft, ArrowRight } from 'iconoir-react'
import { CodeBlock } from '@/components/CodeBlock'
import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'dryLang - Docs',
}

export default function DocsIndex() {
  const docsDir = path.join(process.cwd(), 'docs')
  const filePath = path.join(docsDir, `index.md`)
  
  if (!fs.existsSync(filePath)) {
    return (
      <article className="markdown-body">
        <h1>Welcome to dryLang Documentation</h1>
        <p>Please select a topic from the sidebar to begin.</p>
      </article>
    )
  }

  let content = fs.readFileSync(filePath, 'utf8')
  
  // Strip various formats of Prev/Next pagination at the bottom of files
  content = content.replace(/^.*\[?<?\s*(Prev|Next|Home).*$/gim, '') 
  
  // Strip HTML comments
  content = content.replace(/<!--[\s\S]*?-->/g, '')

  // Strip trailing horizontal rules and whitespace
  content = content.replace(/---\s*$/g, '')
  content = content.trim()

  // Calculate prev/next from DOCS_STRUCTURE
  const flatItems = DOCS_STRUCTURE.flatMap(group => 
    group.items.map(item => ({ ...item, group: group.group }))
  )
  
  // For the home page, current index is 0 (assuming Home is the first item)
  const currentIndex = 0
  const prevItem = null
  const nextItem = flatItems.length > 1 ? flatItems[1] : null

  return (
    <article className="markdown-body" style={{ paddingBottom: '4rem' }}>
      <ReactMarkdown 
        remarkPlugins={[remarkGfm]}
        components={{
          a: ({ node, href, ...props }) => {
            // Strip .md from internal links
            let resolvedHref = href;
            if (resolvedHref && !resolvedHref.startsWith('http')) {
              resolvedHref = resolvedHref.replace(/\.md$/, '');
            }
            return <a href={resolvedHref} {...props} />
          },
          pre: ({children}) => <>{children}</>,
          code: CodeBlock as any
        }}
      >
        {content}
      </ReactMarkdown>

      {/* Dynamic Prev / Next Buttons */}
      <div style={{
        display: 'flex',
        justifyContent: 'space-between',
        marginTop: '4rem',
        paddingTop: '2rem',
        borderTop: '1px solid rgba(255,255,255,0.1)',
        gap: '1rem'
      }}>
        {prevItem ? (
          <Link href={`/docs/${prevItem.slug}`} className="docs-nav-link prev">
            <span style={{ fontSize: '0.85rem', color: '#888', marginBottom: '0.5rem', display: 'flex', alignItems: 'center', gap: '0.25rem' }}>
              <ArrowLeft width={14} height={14} />
              Previous
            </span>
            <span style={{ fontSize: '1.1rem', fontWeight: 500, color: '#fff' }}>{(prevItem as any).title}</span>
          </Link>
        ) : <div style={{ flex: 1 }} />}

        {nextItem ? (
          <Link href={`/docs/${nextItem.slug}`} className="docs-nav-link next">
            <span style={{ fontSize: '0.85rem', color: '#888', marginBottom: '0.5rem', display: 'flex', alignItems: 'center', gap: '0.25rem' }}>
              Next
              <ArrowRight width={14} height={14} />
            </span>
            <span style={{ fontSize: '1.1rem', fontWeight: 500, color: '#fff' }}>{(nextItem as any).title}</span>
          </Link>
        ) : <div style={{ flex: 1 }} />}
      </div>
    </article>
  )
}
