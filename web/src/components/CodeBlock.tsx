"use client";

import React from 'react';

function highlightDryLang(code: string) {
  let html = code
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;');

  const regex = /(\/\/.*)|(".*?")|(\b(?:fn|if|el|elif|lp|pt|cns|rev|brk|ret|imp|len|push|pop|del|map|keys|class|new|this|db|math|op|now|date|arg|env|cmd|dir|req|json|spl|trm|has|q|die)\b)|(\b\d+(?:,\d+)?\b)|(\b(?:tr|fs)\b)|(\b[a-zA-Z_]\w*\b(?=\())|(\?\?)/g;

  return html.replace(regex, (match, p1, p2, p3, p4, p5, p6, p7) => {
    if (p1) return `<span style="color: #6a9955;">${p1}</span>`; // comments
    if (p2) return `<span style="color: #ce9178;">${p2}</span>`; // strings
    if (p3) return `<span style="color: #c586c0;">${p3}</span>`; // keywords & built-ins
    if (p4) return `<span style="color: #b5cea8;">${p4}</span>`; // numbers
    if (p5) return `<span style="color: #569cd6;">${p5}</span>`; // booleans (tr/fs)
    if (p6) return `<span style="color: #dcdcaa;">${p6}</span>`; // function calls
    if (p7) return `<span style="color: #c586c0;">${p7}</span>`; // ?? operator
    return match;
  });
}

function highlightBash(code: string) {
  let html = code
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;');

  const regex = /(#.*)|(^[\$\>]\s)|(\b(?:git|cd|go|npm|node|npx|y|dry)\b)/gm;
  return html.replace(regex, (match, p1, p2, p3) => {
    if (p1) return `<span style="color: #6a9955;">${p1}</span>`;
    if (p2) return `<span style="color: #569cd6;">${p2}</span>`;
    if (p3) return `<span style="color: #dcdcaa;">${p3}</span>`;
    return match;
  });
}

export function CodeBlock({ node, className, children, ...props }: any) {
  const match = /language-(\w+)/.exec(className || '');
  const isInline = !match && !String(children).includes('\n');

  if (isInline) {
    return (
      <code className={className} {...props} style={{
        backgroundColor: 'rgba(255, 255, 255, 0.1)',
        padding: '0.2rem 0.4rem',
        borderRadius: '4px',
        fontSize: '0.9em'
      }}>
        {children}
      </code>
    );
  }

  const language = match ? match[1].toLowerCase() : 'text';

  // If user accidentally put ```javascript for drylang code, or used aliases
  let displayLanguage = language;
  if (['dry', 'drylang', 'rust', 'javascript', 'js', 'ts', 'typescript'].includes(language)) {
    displayLanguage = 'drylang';
  }

  const codeStr = String(children).replace(/\n$/, '');
  let highlightedHtml = codeStr
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;');

  if (displayLanguage === 'drylang') {
    highlightedHtml = highlightDryLang(codeStr);
  } else if (language === 'bash' || language === 'sh') {
    highlightedHtml = highlightBash(codeStr);
  } else {
    // Fallback for completely unknown languages (just use drylang highlighter so it has SOME color)
    highlightedHtml = highlightDryLang(codeStr);
  }

  return (
    <div style={{ position: 'relative', margin: '1.5rem 0', borderRadius: '0', background: '#000000', border: '1px solid rgba(255,255,255,0.1)', overflow: 'hidden' }}>
      <div style={{
        position: 'absolute',
        top: 0,
        right: 0,
        background: 'rgba(255, 255, 255, 0.05)',
        padding: '0.2rem 0.6rem',
        borderRadius: '0',
        fontSize: '0.75rem',
        color: '#aaa',
        fontWeight: 600,
      }}>
        {displayLanguage}
      </div>
      <pre style={{ margin: 0, padding: '1.5rem', overflowX: 'auto', background: 'transparent', border: 'none', borderRadius: 0 }}>
        <code
          dangerouslySetInnerHTML={{ __html: highlightedHtml }}
          style={{
            background: 'transparent',
            padding: 0,
            border: 'none',
            color: '#d4d4d4',
            fontFamily: 'var(--font-geist-mono), monospace',
            fontSize: '0.95rem',
            lineHeight: 1.6,
            fontWeight: 400
          }}
        />
      </pre>
    </div>
  );
}
