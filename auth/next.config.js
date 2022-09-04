const withPlugins = require("next-compose-plugins");
const withPWA = require("next-pwa");

const isDev = process.env.NODE_ENV !== "production";

const nextConfig = {
  env: {
    API_URL: "$PASSWORD_MANAGER_API_WEBSITE_URL",
  },
  pwa: {
    dest: "public",
    disable: isDev,
  },
  reactStrictMode: true,
  future: {
    webpack5: true,
  },
  typescript: {
    ignoreBuildErrors: true,
  },
};

module.exports = withPlugins([], withPWA(nextConfig));
