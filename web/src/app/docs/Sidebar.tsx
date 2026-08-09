'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { DOCS_STRUCTURE } from './config';

export function DocsSidebar() {
  const pathname = usePathname();

  return (
    <>
      {/* Desktop Sidebar */}
      <aside className="docs-sidebar desktop-only">
        {DOCS_STRUCTURE.map((group) => (
          <div key={group.group} className="docs-group">
            <h3 className="docs-group-title">{group.group}</h3>
            <ul className="docs-group-list">
              {group.items.map((item) => {
                const href = item.slug ? `/docs/${item.slug}` : `/docs`;
                const isActive = pathname === href;
                return (
                  <li key={item.slug || 'home'}>
                    <Link href={href} className={`docs-link ${isActive ? 'active' : ''}`}>
                      {item.title}
                    </Link>
                  </li>
                );
              })}
            </ul>
          </div>
        ))}
      </aside>

      {/* Mobile Sidebar (Collapsible) */}
      <details className="mobile-only" style={{ marginBottom: '1rem', cursor: 'pointer', position: 'relative', overflow: 'visible' }}>
        <summary style={{ padding: '1rem 0', fontWeight: 600, borderBottom: '1px solid var(--border-color)', outline: 'none' }}>
          Docs Menu
        </summary>
        <div style={{ position: 'absolute', top: '100%', left: 0, width: '100%', backgroundColor: 'var(--bg-primary)', padding: '1rem', borderBottom: '1px solid var(--border-color)', zIndex: 40, maxHeight: '70vh', overflowY: 'auto' }}>
          {DOCS_STRUCTURE.map((group) => (
            <div key={`mobile-${group.group}`} className="docs-group">
              <h3 className="docs-group-title">{group.group}</h3>
              <ul className="docs-group-list">
                {group.items.map((item) => {
                  const href = item.slug ? `/docs/${item.slug}` : `/docs`;
                  const isActive = pathname === href;
                  return (
                    <li key={item.slug || 'home'}>
                      <Link href={href} className={`docs-link ${isActive ? 'active' : ''}`}>
                        {item.title}
                      </Link>
                    </li>
                  );
                })}
              </ul>
            </div>
          ))}
        </div>
      </details>
    </>

  );
}
