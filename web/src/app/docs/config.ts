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
    group: 'Welcome',
    items: [
      { title: 'Home', slug: '' },
    ],
  },
  {
    group: 'Getting Started',
    items: [
      { title: 'Overview', slug: 'getting-started' },
      { title: 'Installation', slug: 'getting-started/installation' },
      { title: 'CLI Usage', slug: 'getting-started/cli' },
    ],
  },
  {
    group: 'Basics',
    items: [
      { title: 'Overview', slug: 'basics' },
      { title: 'Variables', slug: 'basics/variables' },
      { title: 'Data Types', slug: 'basics/data-types' },
      { title: 'Operators', slug: 'basics/operators' },
      { title: 'Comments', slug: 'basics/comments' },
    ],
  },
  {
    group: 'Control Flow',
    items: [
      { title: 'Overview', slug: 'control-flow' },
      { title: 'Conditionals', slug: 'control-flow/conditionals' },
      { title: 'Loops', slug: 'control-flow/loops' },
      { title: 'Error Handling', slug: 'control-flow/error-handling' },
    ],
  },
  {
    group: 'Functions',
    items: [
      { title: 'Overview', slug: 'functions' },
      { title: 'Definition', slug: 'functions/definition' },
      { title: 'Arrow Functions', slug: 'functions/arrow-functions' },
      { title: 'Scope & Closures', slug: 'functions/scope-closures' },
    ],
  },
  {
    group: 'Data Structures',
    items: [
      { title: 'Overview', slug: 'data-structures' },
      { title: 'Arrays', slug: 'data-structures/arrays' },
      { title: 'Maps', slug: 'data-structures/maps' },
      { title: 'Structs & Classes', slug: 'data-structures/structs' },
    ],
  },
  {
    group: 'Built-ins',
    items: [
      { title: 'Overview', slug: 'built-ins' },
      { title: 'Input & Output', slug: 'built-ins/io' },
      { title: 'Strings', slug: 'built-ins/strings' },
      { title: 'Math', slug: 'built-ins/math' },
      { title: 'Collections', slug: 'built-ins/collections' },
    ],
  },
  {
    group: 'Advanced Features',
    items: [
      { title: 'Overview', slug: 'advanced' },
      { title: 'File System', slug: 'advanced/file-system' },
      { title: 'Networking', slug: 'advanced/networking' },
      { title: 'Database', slug: 'advanced/database' },
      { title: 'System & OS', slug: 'advanced/system' },
    ],
  }
];
