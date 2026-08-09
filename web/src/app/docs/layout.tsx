import React from 'react';
import Link from 'next/link';
import { DocsSidebar } from './Sidebar';

export default function DocsLayout({ children }: { children: React.ReactNode }) {
  return (
    <div style={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>


      {/* Docs Content Area */}
      <div className="container docs-layout">
        <DocsSidebar />
        <main className="docs-main">
          {children}
        </main>
      </div>
    </div>
  );
}
