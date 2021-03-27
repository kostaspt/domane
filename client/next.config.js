const { PHASE_PRODUCTION_BUILD } = require('next/constants');

module.exports = (phase) => {
  const isProd = process.env.NODE_ENV === 'production' || phase === PHASE_PRODUCTION_BUILD;

  return {
    publicRuntimeConfig: {
      API_URL: isProd ? 'https://domane.io/api' : 'http://localhost:4000',
    },
    webpack(config) {
      config.module.rules.push({
        test: /\.svg$/,
        use: ['@svgr/webpack'],
      });

      return config;
    },
  };
};
