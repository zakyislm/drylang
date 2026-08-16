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
      { title: 'CLI & Commands', slug: 'getting-started/cli' },
    ],
  },
  {
    group: 'Basics',
    items: [
      { title: 'Overview', slug: 'basics' },
      { title: 'Comments', slug: 'basics/comments' },
      { title: 'Variables', slug: 'basics/variables' },
      { title: 'Constants', slug: 'basics/constants' },
      { title: 'Data Types', slug: 'basics/data-types' },
      { title: 'Numbers', slug: 'basics/numbers' },
      { title: 'Strings', slug: 'basics/strings' },
      { title: 'String Interpolation', slug: 'basics/string-interpolation' },
      { title: 'Booleans & Unknown', slug: 'basics/booleans-unknown' },
      { title: 'Operators', slug: 'basics/operators' },
      { title: 'Type Checking', slug: 'basics/type-checking' },
    ],
  },
  {
    group: 'Control Flow',
    items: [
      { title: 'Overview', slug: 'control-flow' },
      { title: 'If / Elif / El', slug: 'control-flow/conditionals' },
      { title: 'Switch (on)', slug: 'control-flow/switch-on' },
      { title: 'Loops (lp)', slug: 'control-flow/loops' },
      { title: 'Break & Continue', slug: 'control-flow/break-continue' },
      { title: 'Error Handling', slug: 'control-flow/error-handling' },
    ],
  },
  {
    group: 'Functions',
    items: [
      { title: 'Overview', slug: 'functions' },
      { title: 'Defining Functions', slug: 'functions/definition' },
      { title: 'Arrow Functions', slug: 'functions/arrow-functions' },
      { title: 'Returns (rev)', slug: 'functions/returns' },
      { title: 'Scope & Closures', slug: 'functions/scope-closures' },
    ],
  },
  {
    group: 'Data Structures',
    items: [
      { title: 'Overview', slug: 'data-structures' },
      { title: 'Arrays', slug: 'data-structures/arrays' },
      { title: 'Array Methods', slug: 'data-structures/array-methods' },
      { title: 'Maps (Dictionaries)', slug: 'data-structures/maps' },
      { title: 'Map Methods', slug: 'data-structures/map-methods' },
      { title: 'Structs', slug: 'data-structures/structs' },
      { title: 'Classes (OOP)', slug: 'data-structures/classes' },
      { title: 'Inheritance', slug: 'data-structures/inheritance' },
      { title: 'Private Members', slug: 'data-structures/private-members' },
    ],
  },
  {
    group: 'Async & Concurrency',
    items: [
      { title: 'Introduction (asn)', slug: 'async' },
      { title: 'Multiple Parallel (mul)', slug: 'async/parallel' },
      { title: 'Single Async (uni)', slug: 'async/single' },
      { title: 'Await Sync (awt)', slug: 'async/await' },
      { title: 'Data Streaming (pipe)', slug: 'async/pipe' },
    ],
  },
  {
    group: 'Built-ins',
    items: [
      { title: 'Overview', slug: 'builtins' },
      { title: 'I/O & Network', slug: 'builtins/io' },
      { title: 'String Methods', slug: 'builtins/string-methods' },
      { title: 'Math Methods', slug: 'builtins/math-methods' },
      { title: 'Regular Expressions', slug: 'builtins/regex' },
      { title: 'Crypto Hash', slug: 'builtins/crypto-hash' },
      { title: 'Crypto JWT', slug: 'builtins/crypto-jwt' },
    ],
  },
  {
    group: 'Standard Library',
    items: [
      { title: 'CLI Flags', slug: 'standard-library/cli' },
      { title: 'JSON Parser', slug: 'standard-library/json' },
      { title: 'System (OS, Env)', slug: 'standard-library/system' },
      { title: 'Templates', slug: 'standard-library/templates' },
      { title: 'Testing', slug: 'standard-library/testing' },
      { title: 'Validation', slug: 'standard-library/validation' },
    ],
  },
  {
    group: 'Web & Backend',
    items: [
      { title: 'Overview', slug: 'backend' },
      { title: 'HTTP Request (req)', slug: 'backend/http-request' },
      { title: 'HTTP Server (op)', slug: 'backend/http-server' },
      { title: 'Routing (rt)', slug: 'backend/routing' },
      { title: 'WebSockets (ws)', slug: 'backend/websockets' },
      { title: 'RPC', slug: 'backend/rpc' },
      { title: 'Database Queries', slug: 'backend/database' },
      { title: 'Connection Pooling', slug: 'backend/database-pool' },
    ],
  },
  {
    group: 'State & Jobs',
    items: [
      { title: 'Overview', slug: 'state-jobs' },
      { title: 'Flow Control', slug: 'state-jobs/flow' },
      { title: 'Memory Cache', slug: 'state-jobs/memory-cache' },
      { title: 'Session', slug: 'state-jobs/session' },
      { title: 'Rate Limiting', slug: 'state-jobs/rate-limiting' },
      { title: 'Background Jobs', slug: 'state-jobs/background-jobs' },
      { title: 'Cron Jobs', slug: 'state-jobs/cron-jobs' },
    ],
  },
  {
    group: 'Modules',
    items: [
      { title: 'Overview', slug: 'modules' },
    ],
  },
  {
    group: 'Security',
    items: [
      { title: 'Overview', slug: 'security' },
      { title: 'Sandbox Lockdown', slug: 'security/sandbox' },
      { title: 'Server Hardening', slug: 'security/server-hardening' },
    ],
  }
];
