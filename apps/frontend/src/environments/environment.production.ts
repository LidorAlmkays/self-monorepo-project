import env from './.env';

export const environment = {
  production: true,
  apiBaseUrl: 'http://localhost:8000',
  version: env.npm_package_version,
};
