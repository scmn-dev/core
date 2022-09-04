import React from "react";
import Head from "next/head";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { ToastContainer, toast } from "react-toastify";
import * as yup from "yup";
import Api from "api";
import { useRouter } from "next/router";
import { TextInput } from "@components/text-input";
import { ErrorMsg } from "components/errors";
import { LogoSection } from "components/logo";

const schema = yup.object().shape({
  code: yup
    .string()
    .matches(/^\d+$/, {
      message: "Code can only contain numbers",
    })
    .required(),
});

function CodePage() {
  const router = useRouter();

  const { register, handleSubmit, errors } = useForm({
    resolver: yupResolver(schema),
  });

  const verifyCodeAPI = async ({ code, email }) => {
    return await Api.get(`/auth/verify/` + code + `?email=` + email)
      .then((data) => Promise.resolve(data))
      .catch((err) => {
        toast(<ErrorMsg msg={err.response.data.message} />);

        return Promise.reject(err);
      });
  };

  const onSubmit = ({ code }) => {
    let email = localStorage.getItem("email");

    verifyCodeAPI({ code, email })
      .then(() => {
        router.push("/final-step");
      })
      .catch((err) => console.error(err));
  };

  return (
    <>
      <ToastContainer hideProgressBar />

      <main className="flex flex-col md:flex-row-reverse md:h-screen">
        <Head>
          <title>Verify Code - $PASSWORD_MANAGER_NAME</title>
          <link rel="icon" href="$LOGO_URL" />
        </Head>
        <LogoSection />

        <section className="justify-center px-4 md:px-0 md:flex md:w-2/3">
          <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
            <h2 className="text-lg font-medium md:text-2xl">Enter your code</h2>
            <div className="my-4">
              <form onSubmit={handleSubmit(onSubmit)} method="POST">
                <div className="mb-3">
                  <TextInput
                    className=""
                    label="Please enter the verification code at your email inbox."
                    name="code"
                    placeholder="138452"
                    register={register()}
                    errors={errors.code}
                  />
                </div>
                <div className="mt-6 space-y-2 flex justify-center">
                  <button
                    type="submit"
                    style={{ background: "#1163E6" }}
                    className="button button__md button__primary w-full"
                  >
                    Continue to final step
                  </button>
                </div>
              </form>
            </div>
          </div>
        </section>
      </main>
    </>
  );
}

export default CodePage;
