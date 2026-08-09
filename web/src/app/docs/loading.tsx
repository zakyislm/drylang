import React from 'react';

export default function DocsLoading() {
  return (
    <article className="skeleton-wrapper" style={{ paddingBottom: '4rem' }}>
      {/* H1 Title */}
      <div className="skeleton-box" style={{ height: '3rem', width: '50%', marginBottom: '2rem', marginTop: '0.5rem' }}></div>
      
      {/* Paragraphs */}
      <div className="skeleton-box" style={{ height: '1.1rem', width: '100%', marginBottom: '0.8rem' }}></div>
      <div className="skeleton-box" style={{ height: '1.1rem', width: '95%', marginBottom: '0.8rem' }}></div>
      <div className="skeleton-box" style={{ height: '1.1rem', width: '85%', marginBottom: '2.5rem' }}></div>
      
      {/* H2 Title */}
      <div className="skeleton-box" style={{ height: '2rem', width: '35%', marginBottom: '1.5rem' }}></div>
      
      {/* Code block */}
      <div className="skeleton-box" style={{ height: '220px', width: '100%', borderRadius: '0', marginBottom: '2.5rem' }}></div>

      {/* Paragraphs */}
      <div className="skeleton-box" style={{ height: '1.1rem', width: '90%', marginBottom: '0.8rem' }}></div>
      <div className="skeleton-box" style={{ height: '1.1rem', width: '70%', marginBottom: '2.5rem' }}></div>
      
      {/* List items */}
      <div style={{ display: 'flex', gap: '1rem', marginBottom: '0.8rem', alignItems: 'center' }}>
        <div className="skeleton-box" style={{ height: '8px', width: '8px', borderRadius: '50%' }}></div>
        <div className="skeleton-box" style={{ height: '1.1rem', width: '60%' }}></div>
      </div>
      <div style={{ display: 'flex', gap: '1rem', marginBottom: '0.8rem', alignItems: 'center' }}>
        <div className="skeleton-box" style={{ height: '8px', width: '8px', borderRadius: '50%' }}></div>
        <div className="skeleton-box" style={{ height: '1.1rem', width: '75%' }}></div>
      </div>
      <div style={{ display: 'flex', gap: '1rem', marginBottom: '3rem', alignItems: 'center' }}>
        <div className="skeleton-box" style={{ height: '8px', width: '8px', borderRadius: '50%' }}></div>
        <div className="skeleton-box" style={{ height: '1.1rem', width: '50%' }}></div>
      </div>

      {/* Pagination Buttons */}
      <div style={{ display: 'flex', gap: '1rem', marginTop: '4rem', borderTop: '1px solid var(--border-color)', paddingTop: '2rem' }}>
        <div className="skeleton-box" style={{ flex: 1, height: '90px', borderRadius: '0' }}></div>
        <div className="skeleton-box" style={{ flex: 1, height: '90px', borderRadius: '0' }}></div>
      </div>
    </article>
  )
}
