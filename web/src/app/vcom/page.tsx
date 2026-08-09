import type { Metadata } from 'next'
import VcomClient from './VcomClient'

export const metadata: Metadata = {
  title: 'dryLang - Playground',
}

export default function VcomPage() {
  return <VcomClient />
}
