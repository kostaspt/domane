const { PHASE_PRODUCTION_BUILD } = require('next/constants');

module.exports = (phase) => {
  const isProd = phase === PHASE_PRODUCTION_BUILD;

  return {
    publicRuntimeConfig: {
      API_URL: isProd ? '/api' : 'http://localhost:4000',
    },
  };
};
