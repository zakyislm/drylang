import type { Metadata } from 'next'
import Link from 'next/link'

export const metadata: Metadata = {
  title: 'dryLang - Download',
}

export default function Get() {
  return (
    <main>


      <section className="section">
        <div className="container">
          <h1 className="title-huge" style={{ fontSize: '4rem', marginBottom: '1rem' }}>get drylang.</h1>
          <p className="subtitle" style={{ marginBottom: '4rem' }}>
            the official distribution includes the compiler, bytecode vm, and standard library. choose your platform below.
          </p>

          <div style={{ display: 'flex', flexDirection: 'column', gap: '3rem' }}>
            
            {/* macOS */}
            <div>
              <h2 style={{ fontSize: '1.5rem', marginBottom: '1rem' }}>macos</h2>
              <p style={{ color: 'var(--text-muted)', marginBottom: '1.5rem' }}>
                download and install via the official shell script. requires curl.
              </p>
              <pre className="code-block" style={{ width: 'fit-content', maxWidth: '100%' }}>
                <code>
curl -sSL https://raw.githubusercontent.com/zakyislm/drylang/main/installers/install.sh | bash
                </code>
              </pre>
            </div>

            {/* Linux */}
            <div>
              <h2 style={{ fontSize: '1.5rem', marginBottom: '1rem' }}>linux</h2>
              <p style={{ color: 'var(--text-muted)', marginBottom: '1.5rem' }}>
                download and install via the official shell script. requires curl.
              </p>
              <pre className="code-block" style={{ width: 'fit-content', maxWidth: '100%' }}>
                <code>
curl -sSL https://raw.githubusercontent.com/zakyislm/drylang/main/installers/install.sh | bash
                </code>
              </pre>
            </div>

            {/* Windows */}
            <div>
              <h2 style={{ fontSize: '1.5rem', marginBottom: '1rem' }}>windows</h2>
              <p style={{ color: 'var(--text-muted)', marginBottom: '1.5rem' }}>
                download and install via powershell. adds drylang to your path automatically.
              </p>
              <pre className="code-block" style={{ width: 'fit-content', maxWidth: '100%' }}>
                <code>
iwr https://raw.githubusercontent.com/zakyislm/drylang/main/installers/install.ps1 -useb | iex
                </code>
              </pre>
            </div>

          </div>

          <div style={{ marginTop: '4rem' }}>
            <h2 style={{ fontSize: '1.5rem', marginBottom: '1rem' }}>verify installation</h2>
            <p style={{ color: 'var(--text-muted)', marginBottom: '1.5rem' }}>
              after running the script, verify the installation by checking the version in your terminal.
            </p>
            <pre className="code-block" style={{ width: 'fit-content', maxWidth: '100%' }}>
              <code>y --version</code>
            </pre>
          </div>

        </div>
      </section>
    </main>
  )
}
