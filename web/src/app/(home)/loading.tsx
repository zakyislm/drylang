import React from 'react';

export default function Loading() {
  return (
    <main className="skeleton-wrapper">
      <section className="section">
        <div className="container">
          <div className="skeleton-box" style={{ height: '4rem', width: '70%', marginBottom: '1.5rem', maxWidth: '800px' }}></div>
          <div className="skeleton-box" style={{ height: '1.5rem', width: '90%', marginBottom: '0.5rem', maxWidth: '600px' }}></div>
          <div className="skeleton-box" style={{ height: '1.5rem', width: '75%', maxWidth: '500px' }}></div>
          
          <div style={{ marginTop: '3rem', display: 'flex', gap: '1rem' }}>
             <div className="skeleton-box" style={{ height: '3.5rem', width: '200px' }}></div>
             <div className="skeleton-box" style={{ height: '3.5rem', width: '200px' }}></div>
          </div>
        </div>
      </section>
      
      <section className="section">
        <div className="container">
          <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '4rem', alignItems: 'center' }}>
            <div>
              <div className="skeleton-box" style={{ height: '3rem', width: '60%', marginBottom: '1.5rem' }}></div>
              <div className="skeleton-box" style={{ height: '1.25rem', width: '100%', marginBottom: '0.5rem' }}></div>
              <div className="skeleton-box" style={{ height: '1.25rem', width: '90%', marginBottom: '1.5rem' }}></div>
              <div className="skeleton-box" style={{ height: '1.25rem', width: '85%' }}></div>
            </div>
            
            <div>
               <div className="skeleton-box" style={{ height: '250px', width: '100%', borderRadius: '8px' }}></div>
               <div className="skeleton-box" style={{ marginTop: '1rem', height: '3.5rem', width: '100%' }}></div>
            </div>
          </div>
        </div>
      </section>
      
      <section className="section">
        <div className="container">
          <div className="grid-metrics">
            {[1, 2, 3].map((i) => (
              <div key={i} className="metric-item" style={{ border: '1px solid transparent', background: 'transparent' }}>
                 <div className="skeleton-box" style={{ height: '3.5rem', width: '100px', margin: '0 auto 1rem' }}></div>
                 <div className="skeleton-box" style={{ height: '1.25rem', width: '120px', margin: '0 auto' }}></div>
              </div>
            ))}
          </div>
        </div>
      </section>
      
      <footer className="container" style={{ paddingTop: '4rem', paddingBottom: '4rem', borderTop: '1px solid var(--border-color)', display: 'flex', justifyContent: 'space-between' }}>
         <div className="skeleton-box" style={{ height: '1rem', width: '180px' }}></div>
         <div style={{ display: 'flex', gap: '2rem' }}>
            <div className="skeleton-box" style={{ height: '1rem', width: '40px' }}></div>
            <div className="skeleton-box" style={{ height: '1rem', width: '50px' }}></div>
         </div>
      </footer>
    </main>
  )
}
