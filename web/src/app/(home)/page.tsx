import type { Metadata } from 'next'
import Link from 'next/link'

export const metadata: Metadata = {
  title: 'dryLang - Home',
}

export default function Home() {
  return (
    <main>


      <section className="section">
        <div className="container">
          <h1 className="title-huge">writeless, get more.</h1>
          <p className="subtitle">
            a strict, dynamically-typed bytecode virtual machine. stripped of boilerplate. no syntax exceeds four characters.
          </p>
          <div style={{ marginTop: '3rem', display: 'flex', gap: '1rem' }}>
            <Link href="/get" className="btn-primary">download binary</Link>
            <Link href="/vcom" className="btn-outline">launch compiler</Link>
          </div>
        </div>
      </section>

      <section className="section">
        <div className="container">
          <div className="grid-split">
            <div>
              <h2 style={{ fontSize: '2.5rem', marginBottom: '1.5rem' }}>the manifesto.</h2>
              <p style={{ color: 'var(--text-muted)', fontSize: '1.1rem', marginBottom: '1rem' }}>
                we have traded clarity for ceremony. frameworks breed boilerplate. abstractions multiply complexity.
              </p>
              <p style={{ color: 'var(--text-muted)', fontSize: '1.1rem' }}>
                drylang is the correction. say exactly what you mean. nothing more.
              </p>
            </div>
            
            <div>
              <pre className="code-block">
                <code>
{`fn handler(req) {
    rev "hello, world!"
}
op(8080, handler, "mul", 100)`}
                </code>
              </pre>
              <div style={{ marginTop: '1rem' }}>
                <Link href="/vcom?src=docs#http-server&sc=1" className="btn-outline" style={{ width: '100%', textAlign: 'center' }}>execute code</Link>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section className="section">
        <div className="container">
          <div className="grid-metrics">
            <div className="metric-item">
              <div className="metric-value">5mb</div>
              <div className="metric-label">single binary</div>
            </div>
            <div className="metric-item">
              <div className="metric-value">0</div>
              <div className="metric-label">dependencies</div>
            </div>
            <div className="metric-item">
              <div className="metric-value">38</div>
              <div className="metric-label">core functions</div>
            </div>
          </div>
        </div>
      </section>

      <footer className="container" style={{ paddingTop: '4rem', paddingBottom: '4rem', borderTop: '1px solid var(--border-color)', color: 'var(--text-muted)', fontSize: '0.875rem' }}>
        <div style={{ display: 'flex', justifyContent: 'space-between' }}>
          <div>© 2026 zakyislm. mit licensed.</div>
          <div style={{ display: 'flex', gap: '2rem' }}>
            <Link href="/docs">docs</Link>
            <a href="https://github.com/zakyislm/drylang">github</a>
          </div>
        </div>
      </footer>
    </main>
  )
}
