import fs from 'fs'
import path from 'path'
import { notFound } from 'next/navigation'
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import { DOCS_STRUCTURE } from '../config'
import Link from 'next/link'
import { ArrowLeft, ArrowRight } from 'iconoir-react'
import { CodeBlock } from '@/components/CodeBlock'
import type { Metadata } from 'next'

// Ensure static rendering for known slugs
export function generateStaticParams() {
  const slugs = DOCS_STRUCTURE.flatMap(group => group.items.map(item => ({
    slug: item.slug.split('/')
  }))).filter(item => item.slug.length > 0 && item.slug[0] !== '')
  return slugs
}

export async function generateMetadata({ params }: { params: Promise<{ slug: string[] }> }): Promise<Metadata> {
  const resolvedParams = await params
  const joinedSlug = resolvedParams.slug.join('/')
  const flatItems = DOCS_STRUCTURE.flatMap(group => group.items)
  const item = flatItems.find(i => i.slug === joinedSlug)
  
  if (item && item.title) {
    return {
      title: `dryLang Docs - ${item.title}`,
      description: `Documentation for ${item.title} in dryLang.`,
      openGraph: {
        title: `dryLang Docs - ${item.title}`,
        description: `Documentation for ${item.title} in dryLang.`,
        siteName: 'dryLang',
        type: 'website',
      }
    }
  }
  return {
    title: 'dryLang Docs',
    openGraph: {
      title: 'dryLang Docs',
      siteName: 'dryLang',
      type: 'website',
    }
  }
}

export default async function DocPage({ params }: { params: Promise<{ slug: string[] }> }) {
  const resolvedParams = await params
  const joinedSlug = resolvedParams.slug.join('/')
  const docsDir = path.join(process.cwd(), 'docs')
  
  let filePath = path.join(docsDir, `${joinedSlug}.mdx`)
  
  if (!fs.existsSync(filePath)) {
    // Try checking if it's an index file of a directory
    filePath = path.join(docsDir, joinedSlug, 'index.mdx')
    if (!fs.existsSync(filePath)) {
      notFound()
    }
  }

  let content = fs.readFileSync(filePath, 'utf8')
  
  // Strip various formats of Prev/Next pagination at the bottom of files
  content = content.replace(/^.*\[?<?\s*(Prev|Next|Home).*$/gim, '') 
  
  // Strip HTML comments
  content = content.replace(/<!--[\s\S]*?-->/g, '')

  // Strip trailing horizontal rules and whitespace
  content = content.replace(/---\s*$/g, '')
  content = content.trim()
  
  const flatItems = DOCS_STRUCTURE.flatMap(group => 
    group.items.map(item => ({ ...item, group: group.group }))
  )
  
  const currentIndex = flatItems.findIndex(item => item.slug === joinedSlug)
  const currentItemIndex = currentIndex

  const prevItem = currentItemIndex > 0 ? flatItems[currentItemIndex - 1] : null
  const nextItem = currentItemIndex !== -1 && currentItemIndex < flatItems.length - 1 ? flatItems[currentItemIndex + 1] : null

  return (
    <article className="markdown-body" style={{ paddingBottom: '4rem' }}>
      <ReactMarkdown 
        remarkPlugins={[remarkGfm]}
        components={{
          a: ({ node, href, ...props }) => {
            // Strip .mdx/.md from internal links
            let resolvedHref = href;
            if (resolvedHref && !resolvedHref.startsWith('http')) {
              resolvedHref = resolvedHref.replace(/\.mdx?$/, '');
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
            <span style={{ fontSize: '1.1rem', fontWeight: 500, color: '#fff' }}>{prevItem.title}</span>
          </Link>
        ) : <div style={{ flex: 1 }} />}

        {nextItem ? (
          <Link href={`/docs/${nextItem.slug}`} className="docs-nav-link next">
            <span style={{ fontSize: '0.85rem', color: '#888', marginBottom: '0.5rem', display: 'flex', alignItems: 'center', gap: '0.25rem' }}>
              Next
              <ArrowRight width={14} height={14} />
            </span>
            <span style={{ fontSize: '1.1rem', fontWeight: 500, color: '#fff' }}>{nextItem.title}</span>
          </Link>
        ) : <div style={{ flex: 1 }} />}
      </div>
    </article>
  )
}
