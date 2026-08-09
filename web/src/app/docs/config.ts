export interface DocItem {
  title: string;
  slug: string;
}

export interface DocGroup {
  group: string;
  items: DocItem[];
}

export const DOCS_STRUCTURE: DocGroup[] = [
  {
    group: 'Getting Started',
    items: [
      { title: 'Home', slug: '' },
      { title: 'Getting Started', slug: 'getting-started' },
      { title: 'Setup', slug: 'cli' },
    ],
  },
  {
    group: 'Language Basics',
    items: [
      { title: 'Variables', slug: 'variables' },
      { title: 'Types', slug: 'types' },
      { title: 'Operators', slug: 'operators' },
      { title: 'Comments', slug: 'comments' },
      { title: 'Strings', slug: 'strings' },
    ],
  },
  {
    group: 'Control Flow',
    items: [
      { title: 'Control Flow', slug: 'control-flow' },
      { title: 'Loops', slug: 'loops' },
    ],
  },
  {
    group: 'Advanced Concepts',
    items: [
      { title: 'Functions', slug: 'functions' },
      { title: 'Structs', slug: 'structs' },
      { title: 'Collections', slug: 'collections' },
      { title: 'Modules', slug: 'modules' },
      { title: 'HTTP Server', slug: 'http-server' },
      { title: 'Database', slug: 'database' },
    ],
  },
  {
    group: 'Reference',
    items: [
      { title: 'Built-ins', slug: 'builtins' },
      { title: 'Error Handling', slug: 'error-handling' },
      { title: 'Errors', slug: 'errors' },
    ],
  },
  {
    group: 'Examples',
    items: [
      { "title": "Hello", "slug": "ex-01-hello" },
      { "title": "Variables", "slug": "ex-02-variables" },
      { "title": "Functions", "slug": "ex-03-functions" },
      { "title": "Control Flow", "slug": "ex-04-control-flow" },
      { "title": "Structs", "slug": "ex-05-structs" },
      { "title": "Arrays Maps", "slug": "ex-06-arrays-maps" },
      { "title": "Try Err", "slug": "ex-07-try-err" },
      { "title": "Imports", "slug": "ex-08-imports" },
      { "title": "Math", "slug": "ex-09-math" },
      { "title": "Strings", "slug": "ex-10-strings" },
      { "title": "Http Server", "slug": "ex-11-http-server" },
      { "title": "Database", "slug": "ex-12-database" },
      { "title": "File Io", "slug": "ex-13-file-io" }
    ],
  },
  {
    group: 'Templates',
    items: [
      { "title": "Automation", "slug": "tpl-automation" },
      { "title": "Cli Tool", "slug": "tpl-cli-tool" },
      { "title": "Crud", "slug": "tpl-crud" },
      { "title": "Fetch Json", "slug": "tpl-fetch-json" },
      { "title": "File Server", "slug": "tpl-file-server" },
      { "title": "Hello", "slug": "tpl-hello" },
      { "title": "Html Render", "slug": "tpl-html-render" },
      { "title": "Rest Api", "slug": "tpl-rest-api" }
    ],
  }
];
