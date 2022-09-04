import React, { useEffect } from "react";
import Head from "next/head";
import { useRouter } from "next/router";
import { LogoSection } from "components/logo";

export default function Page() {
  const router = useRouter();

  // useEffect(() => {
  //   setTimeout(() => {
  //     router.push("/");
  //   }, 3500);
  // }, []);

  return (
    <>
      <main className="flex flex-col md:flex-row-reverse md:h-screen">
        <Head>
          <title>Done - $PASSWORD_MANAGER_NAME</title>
          <link rel="icon" href="$LOGO_URL" />
        </Head>
        <LogoSection />

        <section className="justify-center px-4 md:px-0 md:flex md:w-2/3">
          <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
            <div className="my-4">
              <p>
                Great, when the process of changing your master password is completed,
                we will send you an email.
                <br />
                <br />
                Have a nice day 👊.
              </p>
            </div>
          </div>
        </section>
      </main>
    </>
  );
}
