import React from 'react';

export default function VcomLoading() {
  return (
    <div className="skeleton-wrapper" style={{ display: 'flex', flexDirection: 'column', height: '100vh', backgroundColor: 'var(--bg-primary)' }}>
      <div style={{ display: 'flex', flex: 1, overflow: 'hidden' }}>
        <div style={{ flex: 1, display: 'flex', flexDirection: 'column', backgroundColor: 'var(--bg-secondary)' }}>
          {/* File Header & Preset Dropdown */}
          <div style={{
            backgroundColor: 'var(--bg-primary)',
            borderBottom: '1px solid var(--border-color)',
            height: '61px', // approx height of the header
          }}>
            <div className="container" style={{ padding: '1rem 2rem', display: 'flex', justifyContent: 'flex-end', alignItems: 'center', height: '100%' }}>
              <div style={{ display: 'flex', gap: '1rem', alignItems: 'center' }}>
                <div className="skeleton-box" style={{ width: '150px', height: '34px', borderRadius: '0' }}></div>
                <div className="skeleton-box" style={{ width: '40px', height: '34px', borderRadius: '0' }}></div>
              </div>
            </div>
          </div>
          
          {/* Main Editor Area */}
          <div style={{ flex: 1, position: 'relative' }}>
            <div className="container" style={{ padding: '2rem' }}>
              <div className="skeleton-box" style={{ height: '1.2rem', width: '40%', marginBottom: '0.8rem' }}></div>
              <div className="skeleton-box" style={{ height: '1.2rem', width: '25%', marginBottom: '0.8rem' }}></div>
              <div className="skeleton-box" style={{ height: '1.2rem', width: '50%', marginBottom: '0.8rem' }}></div>
              <div className="skeleton-box" style={{ height: '1.2rem', width: '35%' }}></div>
            </div>
          </div>

          {/* Terminal Area */}
          <div style={{ flexShrink: 0, height: '300px', borderTop: '1px solid var(--border-color)', backgroundColor: '#0a0a0a', padding: '1rem' }}>
            <div className="skeleton-box" style={{ height: '1rem', width: '200px', marginBottom: '0.8rem' }}></div>
            <div className="skeleton-box" style={{ height: '1rem', width: '300px' }}></div>
          </div>
        </div>
      </div>
    </div>
  )
}
