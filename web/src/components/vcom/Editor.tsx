import React, { useEffect, useRef } from 'react';
import Editor, { Monaco } from '@monaco-editor/react';

interface VcomEditorProps {
  code: string;
  onChange: (value: string | undefined) => void;
}

export function VcomEditor({ code, onChange }: VcomEditorProps) {
  
  const handleEditorWillMount = (monaco: Monaco) => {
    // Register DryLang language
    monaco.languages.register({ id: 'drylang' });

    // Language configuration for auto-closing brackets
    monaco.languages.setLanguageConfiguration('drylang', {
      autoClosingPairs: [
        { open: '{', close: '}' },
        { open: '[', close: ']' },
        { open: '(', close: ')' },
        { open: '"', close: '"' },
      ],
      surroundingPairs: [
        { open: '{', close: '}' },
        { open: '[', close: ']' },
        { open: '(', close: ')' },
        { open: '"', close: '"' },
      ],
      brackets: [
        ['{', '}'],
        ['[', ']'],
        ['(', ')'],
      ],
    });

    // Define syntax highlighter rules
    monaco.languages.setMonarchTokensProvider('drylang', {
      keywords: [
        'fn', 'rev', 'if', 'el', 'elif', 'pt', 'lp', 'br', 'con', 'err', 'try', 'cns', 'on', 'done', 'asn', 'awt', 'pv', 'use', 'unknown', 'in', 't', 'f'
      ],
      operators: [
        '=', '>', '<', '==', '<=', '>=', '!=', '+', '-', '*', '/', '%'
      ],
      tokenizer: {
        root: [
          // Comments: . comment .
          [/^\s*\..*\.\s*$/, 'comment'],
          [/\..*\./, 'comment'],
          
          // Keywords
          [/[a-zA-Z_]\w*/, {
            cases: {
              '@keywords': 'keyword',
              '@default': 'identifier'
            }
          }],

          // Numbers
          [/[0-9]+/, 'number'],

          // Strings
          [/"([^"\\]|\\.)*$/, 'string.invalid'],
          [/"/, 'string', '@string'],
        ],
        string: [
          [/[^\\"]+/, 'string'],
          [/\\./, 'string.escape'],
          [/"/, 'string', '@pop']
        ]
      }
    });

    // Define custom theme
    monaco.editor.defineTheme('drylang-dark', {
      base: 'vs-dark',
      inherit: true,
      rules: [
        { token: 'keyword', foreground: '569cd6' },
        { token: 'identifier', foreground: '9cdcfe' },
        { token: 'number', foreground: 'b5cea8' },
        { token: 'string', foreground: 'ce9178' },
        { token: 'comment', foreground: '6a9955' },
      ],
      colors: {
        'editor.background': '#0f111a',
        
        // Find widget
        'editorWidget.background': '#030712',
        'editorWidget.border': '#1E293B',
        'editorWidget.foreground': '#F3F4F6',
        'input.background': '#090E17',
        'input.foreground': '#F3F4F6',
        'input.border': '#1E293B',
        'widget.shadow': '#000000',
        
        // Suggest widget (Autocomplete)
        'editorSuggestWidget.background': '#030712',
        'editorSuggestWidget.border': '#1E293B',
        'editorSuggestWidget.foreground': '#F3F4F6',
        'editorSuggestWidget.highlightForeground': '#3B82F6',
        'editorSuggestWidget.selectedBackground': '#1E293B',

        // Hover widget
        'editorHoverWidget.background': '#030712',
        'editorHoverWidget.border': '#1E293B',
        
        // Selections
        'editor.selectionBackground': '#1E293B',
        'editor.inactiveSelectionBackground': '#0F172A',
      }
    });

    // Auto-complete provider
    monaco.languages.registerCompletionItemProvider('drylang', {
      provideCompletionItems: (model: any, position: any) => {
        const suggestions = [
          // Snippets for blocks
          {
            label: 'fn',
            kind: monaco.languages.CompletionItemKind.Snippet,
            insertText: 'fn ${1:name}(${2:args}) {\n\t$0\n}',
            insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
            documentation: 'Function definition',
            range: undefined as any
          },
          {
            label: 'lp',
            kind: monaco.languages.CompletionItemKind.Snippet,
            insertText: 'lp ${1:condition} {\n\t$0\n}',
            insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
            documentation: 'Universal loop',
            range: undefined as any
          },
          {
            label: 'if',
            kind: monaco.languages.CompletionItemKind.Snippet,
            insertText: 'if ${1:condition} {\n\t$0\n}',
            insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
            documentation: 'If statement',
            range: undefined as any
          },
          // Keywords (excluding the ones we made snippets for)
          ...['rev', 'el', 'elif', 'pt', 'br', 'con', 'err', 'try', 'cns', 'use', 'on', 'done', 'asn', 'awt', 'pv', 'in'].map(k => ({
            label: k,
            kind: monaco.languages.CompletionItemKind.Keyword,
            insertText: k,
            range: undefined as any
          })),
          // Built-ins
          ...['len', 'get', 'add', 'num', 'str', 'abs', 'min', 'max', 'rnd', 'cap', 'low', 'trm', 'spl', 'j', 'mod', 'has', 'sort', 'pop', 'rm', 'key', 'val', 'ran', 'q'].map(b => ({
            label: b,
            kind: monaco.languages.CompletionItemKind.Function,
            insertText: b + '($1)',
            insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
            range: undefined as any
          }))
        ];
        return { suggestions };
      }
    });
  };

  return (
    <div style={{ flex: 1, width: '100%', height: '100%' }}>
      <Editor
        height="100%"
        language="drylang"
        theme="drylang-dark"
        value={code}
        onChange={onChange}
        beforeMount={handleEditorWillMount}
        options={{
          minimap: { enabled: false },
          fontSize: 14,
          fontFamily: 'Consolas, "Courier New", monospace',
          wordWrap: 'on',
          scrollBeyondLastLine: false,
          automaticLayout: true,
          lineDecorationsWidth: 32,
          padding: { top: 16 },
        }}
      />
    </div>
  );
}
