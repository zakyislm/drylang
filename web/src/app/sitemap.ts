import { MetadataRoute } from 'next'
import { DOCS_STRUCTURE } from './docs/config'

export default function sitemap(): MetadataRoute.Sitemap {
  const baseUrl = 'https://drylang.jeki.me'
  
  const routes = [
    '',
    '/get',
    '/vcom',
    '/docs',
  ].map((route) => ({
    url: `${baseUrl}${route}`,
    lastModified: new Date(),
    changeFrequency: 'daily' as const,
    priority: route === '' ? 1 : 0.8,
  }))

  const docRoutes = DOCS_STRUCTURE.flatMap(group => 
    group.items.filter(item => item.slug).map(item => ({
      url: `${baseUrl}/docs/${item.slug}`,
      lastModified: new Date(),
      changeFrequency: 'weekly' as const,
      priority: 0.6,
    }))
  )

  return [...routes, ...docRoutes]
}
