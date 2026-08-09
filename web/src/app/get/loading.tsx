import React from 'react';

export default function GetLoading() {
  return (
    <main className="skeleton-wrapper">
      <section className="section">
        <div className="container">
          {/* Title */}
          <div className="skeleton-box" style={{ height: '4rem', width: '30%', marginBottom: '1rem', maxWidth: '400px' }}></div>
          
          {/* Subtitle */}
          <div className="skeleton-box" style={{ height: '1.5rem', width: '80%', marginBottom: '4rem', maxWidth: '800px' }}></div>

          <div style={{ display: 'flex', flexDirection: 'column', gap: '3rem' }}>
            
            {/* macOS */}
            <div>
              <div className="skeleton-box" style={{ height: '1.8rem', width: '100px', marginBottom: '1rem' }}></div>
              <div className="skeleton-box" style={{ height: '1.1rem', width: '50%', marginBottom: '1.5rem' }}></div>
              <div className="skeleton-box" style={{ height: '4rem', width: '100%', maxWidth: '800px' }}></div>
            </div>

            {/* Linux */}
            <div>
              <div className="skeleton-box" style={{ height: '1.8rem', width: '80px', marginBottom: '1rem' }}></div>
              <div className="skeleton-box" style={{ height: '1.1rem', width: '50%', marginBottom: '1.5rem' }}></div>
              <div className="skeleton-box" style={{ height: '4rem', width: '100%', maxWidth: '800px' }}></div>
            </div>

            {/* Windows */}
            <div>
              <div className="skeleton-box" style={{ height: '1.8rem', width: '120px', marginBottom: '1rem' }}></div>
              <div className="skeleton-box" style={{ height: '1.1rem', width: '55%', marginBottom: '1.5rem' }}></div>
              <div className="skeleton-box" style={{ height: '4rem', width: '100%', maxWidth: '800px' }}></div>
            </div>

          </div>

          <div style={{ marginTop: '4rem' }}>
            <div className="skeleton-box" style={{ height: '1.8rem', width: '250px', marginBottom: '1rem' }}></div>
            <div className="skeleton-box" style={{ height: '1.1rem', width: '60%', marginBottom: '1.5rem' }}></div>
            <div className="skeleton-box" style={{ height: '4rem', width: '200px' }}></div>
          </div>

        </div>
      </section>
    </main>
  )
}
