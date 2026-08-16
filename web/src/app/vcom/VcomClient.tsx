"use client"

import React, { useState, useEffect } from 'react'
import { decompressFromEncodedURIComponent } from 'lz-string'

import Link from 'next/link'
import { FileNode } from '@/components/vcom/types'
import { VcomEditor } from '@/components/vcom/Editor'
import { Terminal } from '@/components/vcom/Terminal'
import { Play } from 'iconoir-react'

const PRESETS = [
  { id: '1', name: 'hello_world', content: '// Simple hello world\npt("hello world!")' },
  { id: '2', name: 'variables', content: '// Variables and constants\ncns MAX = 100\nscore = 50\npt("Score: ${score}")' },
  { id: '3', name: 'functions', content: '// Functions use fn and rev\nfn add(a, b) {\n  rev a + b\n}\n\npt("Result: ${add(5, 7)}")' },
  { id: '4', name: 'loops', content: '// Loops with lp\nx = 0\nlp x < 3 {\n  pt("Iteration ${x}")\n  x = x + 1\n}' },
  { id: '5', name: 'conditionals', content: '// Conditionals\nage = 20\nif age >= 18 {\n  pt("Adult")\n} el {\n  pt("Minor")\n}' },
  { id: '6', name: 'error_handling', content: '// Try-err blocks\ntry {\n  pt("Attempting risky operation...")\n  err "Simulation failed!"\n} err(e) {\n  pt("Caught: ${e}")\n}' },
]

export default function VcomPage() {
  const [activePresetId, setActivePresetId] = useState<string>('1')
  const [content, setContent] = useState<string>(PRESETS[0].content)

  const [output, setOutput] = useState("")
  const [isWasmReady, setIsWasmReady] = useState(false)

  const hasInitialized = React.useRef(false);

  // Try to load code from URL first, then fallback to local storage
  useEffect(() => {
    if (hasInitialized.current) return;
    hasInitialized.current = true;

    const params = new URLSearchParams(window.location.search);
    const cParam = params.get('c');
    if (cParam) {
      const decompressed = decompressFromEncodedURIComponent(cParam);
      if (decompressed) {
        setContent(decompressed);
        setActivePresetId(''); // Set to empty for custom docs snippet
        // Clean URL to prevent staying on huge query param string
        setTimeout(() => {
          window.history.replaceState({}, '', '/vcom');
        }, 100);
        return;
      }
    }

    const saved = localStorage.getItem('drylang-vcom-code')
    if (saved) {
      setContent(saved)
    }
  }, [])

  useEffect(() => {
    localStorage.setItem('drylang-vcom-code', content)
  }, [content])

  useEffect(() => {
    const loadWasm = async () => {
      if (typeof window === 'undefined' || !(window as any).Go) {
        setTimeout(loadWasm, 100)
        return
      }
      try {
        const go = new (window as any).Go()
        const result = await WebAssembly.instantiateStreaming(fetch('/drylang.wasm'), go.importObject)
        go.run(result.instance)
        setIsWasmReady(true)
      } catch (err) {
        console.error("Failed to load dryLang Wasm:", err)
        setOutput(prev => prev + "bash: drylang: command not found (wasm load failed)\n")
      }
    }
    loadWasm()
  }, [])

  const handleSelectPreset = (id: string) => {
    setActivePresetId(id)
    const p = PRESETS.find(x => x.id === id)
    if (p) {
      setContent(p.content)
      // Clear output on preset change
      setOutput("")
    }
  }

  const handleCommand = (cmd: string) => {
    const trimmed = cmd.trim()
    if (!trimmed) {
      setOutput(prev => prev + "\n")
      return
    }

    if (trimmed === 'clear') {
      setOutput("")
      return
    }

    // if it's the hidden run command triggered by the button
    if (trimmed === 'drylang idx.y') {
      const historyAddition = `> running idx.y\n`

      if (!isWasmReady || typeof (window as any).rundryLang !== 'function') {
        setOutput(prev => prev + historyAddition + `bash: drylang: command not found (compiler loading)\n`)
        return
      }

      let captured = ""
      const originalLog = console.log
      const originalError = console.error

      console.log = (...a) => { captured += a.join(" ") + "\n" }
      console.error = (...a) => { captured += a.join(" ") + "\n" }

      try {
        const errStr = (window as any).rundryLang(content)
        if (errStr) console.error(errStr)
      } catch (err) {
        console.error(String(err))
      } finally {
        console.log = originalLog
        console.error = originalError
        if (captured && !captured.endsWith("\n")) captured += "\n"
        setOutput(prev => prev + historyAddition + captured)
      }
      return
    }

    // Anything else typed directly in the terminal gets echoed (standard input styling)
    setOutput(prev => prev + trimmed + '\n')
  }

  return (
    <div style={{ display: 'flex', flexDirection: 'column', height: '100vh', backgroundColor: 'var(--bg-primary)', position: 'relative' }}>


      {/* Main Two-Column Layout */}
      <div style={{ display: 'flex', flex: 1, overflow: 'hidden' }}>

        {/* Main Editor & Terminal Area */}
        <div style={{ flex: 1, display: 'flex', flexDirection: 'column', backgroundColor: 'var(--bg-secondary)' }}>

          {/* File Header & Preset Dropdown */}
          <div style={{
            backgroundColor: 'var(--bg-primary)',
            borderBottom: '1px solid var(--border-color)',
          }}>
            <div className="container" style={{ padding: '1rem 2rem', display: 'flex', justifyContent: 'flex-end', alignItems: 'center' }}>
              <div style={{ display: 'flex', gap: '1rem', alignItems: 'center' }}>
                <select
                  value={activePresetId}
                  onChange={(e) => handleSelectPreset(e.target.value)}
                  style={{
                    backgroundColor: 'var(--bg-secondary)',
                    border: '1px solid var(--border-color)',
                    color: 'var(--text-primary)',
                    padding: '0.5rem 1rem',
                    fontFamily: 'monospace',
                    fontSize: '0.875rem',
                    outline: 'none',
                    cursor: 'pointer'
                  }}
                >
                  <option value="" disabled style={{ display: 'none' }}>Custom Code</option>
                  {PRESETS.map(p => (
                    <option key={p.id} value={p.id}>{p.name}</option>
                  ))}
                </select>
                <button
                  onClick={() => handleCommand('drylang idx.y')}
                  style={{
                    backgroundColor: 'var(--text-primary)',
                    color: 'var(--bg-primary)',
                    border: 'none',
                    padding: '0.5rem 1rem',
                    cursor: 'pointer',
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center'
                  }}
                >
                  <Play width={16} height={16} />
                </button>
              </div>
            </div>
          </div>
          <div style={{ flex: 1, position: 'relative', minHeight: 0 }}>
            <div className="container" style={{ position: 'absolute', top: 0, bottom: 0, left: 0, right: 0 }}>
              <VcomEditor code={content} onChange={(c) => setContent(c || '')} />
            </div>
          </div>

            <div style={{ flexShrink: 0 }}>
              <Terminal
                output={output}
                isWasmReady={isWasmReady}
                onClear={() => setOutput("")}
                onCommand={handleCommand}
              />
            </div>
          </div>
      </div>
    </div>
  )
}
