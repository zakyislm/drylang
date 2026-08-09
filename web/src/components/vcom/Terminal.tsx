import React from 'react';

interface TerminalProps {
  output: string;
  isWasmReady: boolean;
  onClear: () => void;
  onCommand: (cmd: string) => void;
}

export function Terminal({ output, isWasmReady, onClear, onCommand }: TerminalProps) {
  const [input, setInput] = React.useState('');
  const inputRef = React.useRef<HTMLInputElement>(null);

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter') {
      onCommand(input);
      setInput('');
    }
  };

  return (
    <div
      onClick={() => inputRef.current?.focus()}
      style={{
        backgroundColor: 'var(--bg-secondary)',
        borderTop: '1px solid var(--border-color)',
        cursor: 'text',
        fontFamily: 'monospace',
        fontSize: '0.9rem',
        color: 'var(--text-primary)',
        height: '30vh',
        overflowY: 'auto',
        display: 'flex',
        flexDirection: 'column',
      }}>
      
      <div className="container" style={{ flex: 1, display: 'flex', flexDirection: 'column', padding: '1rem 2rem' }}>
        <div style={{
        fontSize: '0.75rem',
        color: 'var(--text-muted)',
        textTransform: 'uppercase',
        letterSpacing: '0.1em',
        marginBottom: '1rem',
        userSelect: 'none',
        fontWeight: 'bold'
      }}>
        Terminal
      </div>

      <div style={{
        whiteSpace: 'pre-wrap',
        marginBottom: '0.5rem'
      }}>
        {output && <div>{output}</div>}
      </div>

      <div style={{ display: 'flex', alignItems: 'center' }}>
        <input
          ref={inputRef}
          value={input}
          onChange={e => setInput(e.target.value)}
          onKeyDown={handleKeyDown}
          style={{
            background: 'transparent', 
            border: 'none', 
            color: 'var(--text-primary)',
            fontFamily: 'inherit', 
            fontSize: 'inherit', 
            outline: 'none', 
            flex: 1, 
          }}
          placeholder="Enter standard input..."
          spellCheck={false}
          autoFocus
        />
      </div>
      </div>
    </div>
  );
}
