"use client"

import Link from 'next/link';
import { useState } from 'react';

export function Navbar() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div style={{ position: 'sticky', top: 0, zIndex: 50, backgroundColor: 'var(--bg-primary)', borderBottom: '1px solid var(--border-color)' }}>
      <div className="container">
        <nav className="nav">
          <Link href="/" className="nav-logo" style={{ textDecoration: 'none' }}>
            dry<span style={{ color: '#004aad' }}>Lang</span>
          </Link>
          
          {/* Hamburger Icon */}
          <button 
            className="mobile-only" 
            onClick={() => setIsOpen(!isOpen)}
            style={{ background: 'none', border: 'none', color: 'var(--text-primary)', cursor: 'pointer', padding: '0.5rem' }}
          >
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              {isOpen ? (
                <>
                  <line x1="18" y1="6" x2="6" y2="18"></line>
                  <line x1="6" y1="6" x2="18" y2="18"></line>
                </>
              ) : (
                <>
                  <line x1="3" y1="12" x2="21" y2="12"></line>
                  <line x1="3" y1="6" x2="21" y2="6"></line>
                  <line x1="3" y1="18" x2="21" y2="18"></line>
                </>
              )}
            </svg>
          </button>

          <div className={`nav-links ${isOpen ? 'open' : ''}`}>
            <Link href="/docs" onClick={() => setIsOpen(false)}>docs</Link>
            <Link href="/vcom" onClick={() => setIsOpen(false)}>playground</Link>
            <Link href="/get" onClick={() => setIsOpen(false)}>download</Link>
            <a href="https://github.com/zakyislm/drylang" target="_blank" rel="noopener noreferrer" onClick={() => setIsOpen(false)}>github</a>
          </div>
        </nav>
      </div>
    </div>
  );
}
