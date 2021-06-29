const { moduleFileExtensions } = require('jest-config').defaults;
const { pathsToModuleNameMapper } = require('ts-jest/utils');
const { compilerOptions } = require('./tsconfig');

module.exports = {
  bail: true,
  moduleFileExtensions: [...moduleFileExtensions, 'ts', 'tsx'],
  moduleNameMapper: pathsToModuleNameMapper(compilerOptions.paths, { prefix: '<rootDir>/src' }),
  preset: 'ts-jest',
  roots: ['src'],
  setupFilesAfterEnv: ['<rootDir>/jest-setup.ts'],
  testEnvironment: 'jsdom',
  testMatch: ['<rootDir>/src/**/?(*.)test.{ts,tsx}'],
  transform: {
    '\\.svg$': '<rootDir>/test/__mocks__/svgrMock.js',
  },
  verbose: true,
  globals: {
    'ts-jest': {
      tsconfig: {
        jsx: 'react-jsx',
      },
    },
  },
};
