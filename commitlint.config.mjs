export default {
  parserPreset: 'conventional-changelog-atom',
  formatter: '@commitlint/format',
  rules: {
    "header-max-length": [2, "always", 120],
    'subject-empty': [2, 'never'],
    'subject-full-stop': [2, 'never', '.'],
    'subject-max-length': [2, 'always', 120],
    'subject-min-length': [2, 'always', 10],
    'type-enum': [2, 'always', ['ci', 'docs', 'feature', 'fix', 'improvement', 'perf', 'refactor', 'revert', 'style', 'test', 'unit-test', 'build']],
    'type-empty': [2, 'never'],
    'scope-empty': [2, 'never'],
    'scope-min-length': [2, 'always', 3],
    'body-min-length': [2, 'always', 15],
    'body-max-line-length': [2, 'always', 120],
    'body-leading-blank': [2, 'always'],
  },
  parserPreset: {
    parserOpts: {
      noteKeywords: ['\\[\\d+\\]']
    }
  }
};
