const { PHASE_PRODUCTION_BUILD } = require('next/constants');

module.exports = (phase) => {
  const isProd = phase === PHASE_PRODUCTION_BUILD;

  return {
    publicRuntimeConfig: {
      API_URL: isProd ? 'https://domane.io/api' : 'http://localhost:4000',
    },
  };
};
