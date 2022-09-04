import React from "react";
import Head from "next/head";
import "../styles/app.css";
import SiteConfig from "../site.config";

export default function MyApp({ Component, pageProps }) {
  return (
    <>
      <Head>
        {/* base */}
        <meta charSet="utf-8" />
        <meta httpEquiv="x-ua-compatible" content="ie=edge" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <meta name="description" content={SiteConfig.description} />

        {/* facebook */}
        <meta property="og:url" content={SiteConfig.siteUrl} />
        <meta property="og:type" content="website" />
        <meta property="og:title" content={SiteConfig.title} />
        <meta property="og:description" content={SiteConfig.description} />

        {/* twitter */}
        <meta name="twitter:card" content="summary" />
        <meta name="twitter:url" content={SiteConfig.siteUrl} />
        <meta name="twitter:title" content={SiteConfig.title} />
        <meta name="twitter:description" content={SiteConfig.description} />
      </Head>

      <Component {...pageProps} />
    </>
  );
}
