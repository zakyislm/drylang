import './globals.css'
import type { Metadata } from 'next'
import { Navbar } from '@/components/Navbar'

export const metadata: Metadata = {
  title: {
    template: '%s',
    default: 'dryLang - Home',
  },
  description: 'dryLang is a minimalist, dynamically-typed programming language compiled to bytecode. No fluff. Just signal, no noise.',
  keywords: 'dry programming language, drylang, dry language dry lang, drylanguage, dryprogramminglanguage',
  authors: [{ name: 'zakyislm' }],
  icons: {
    icon: '/favicon.svg',
  },
  verification: {
    google: 'NV2TxD_Kt7HqT7tik_TgpEqtuAzQ6bt7n-NaZV8t21A',
  },
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <head>
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link rel="preconnect" href="https://fonts.gstatic.com" crossOrigin="anonymous" />
        <link href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;600&family=Lora:ital,wght@0,400;0,500;0,600;0,700;1,400&display=swap" rel="stylesheet" />
        <script src="/wasm_exec.js" async></script>
      </head>
      <body>
        <div className="app-root">
          <Navbar />
          {children}
        </div>
      </body>
    </html>
  )
}
