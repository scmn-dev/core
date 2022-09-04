import Document, { Html, Head, Main, NextScript } from "next/document";
import React from "react";
import SiteConfig from "site.config";

class MyDocument extends Document {
  static async getInitialProps(ctx) {
    const initialProps = await Document.getInitialProps(ctx);
    return { ...initialProps };
  }

  render() {
    return (
      <Html lang={SiteConfig.lang}>
        <Head>
          <link
            rel="apple-touch-icon"
            href="$LOGO_URL"
          />

          <link
            rel="icon"
            type="image/svg"
            href="$LOGO_URL"
          />

          <meta name="msapplication-TileColor" content="#ffffff" />
          <meta name="theme-color" content="#ffffff" />

          <script src="https://cdn.paddle.com/paddle/paddle.js"></script>
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}

export default MyDocument;
