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
  email: yup.string().email().required(),
});

function ForgotPasswordPage() {
  const router = useRouter();

  const { register, handleSubmit, errors } = useForm({
    resolver: yupResolver(schema),
  });

  const sendEmailAPI = async ({ email }) => {
    return await Api.post(`/auth/fmp-send`, {
      email,
    })
      .then((data) => Promise.resolve(data))
      .catch((err) => {
        toast(<ErrorMsg msg={err.response.data.message} />);

        return Promise.reject(err);
      });
  };

  const onSubmit = ({ email }) => {
    localStorage.setItem("email", email);

    sendEmailAPI({ email })
      .then(() => {
        router.push("/forgot-password/verify");
      })
      .catch((err) => console.error(err));
  };

  return (
    <>
      <ToastContainer hideProgressBar />

      <main className="flex flex-col md:flex-row-reverse md:h-screen">
        <Head>
          <title>Reset your master password - $PASSWORD_MANAGER_NAME</title>
          <link rel="icon" href="$LOGO_URL" />
        </Head>
        <LogoSection />

        <section className="justify-center px-4 md:px-0 md:flex md:w-2/3">
          <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
            <h2 className="text-lg font-medium md:text-2xl">
              Reset your master password
            </h2>
            <p className="text-sm pt-2 text-gray-400">
              Enter your email address and we will send you instructions to
              reset your master password.
            </p>
            <div className="my-4">
              <form onSubmit={handleSubmit(onSubmit)} method="POST">
                <div className="mb-3">
                  <label
                    htmlFor="email"
                    className="block text-sm font-medium text-neutral-400"
                  >
                    Email address
                  </label>
                  <TextInput
                    name="email"
                    type="email"
                    placeholder="Email"
                    register={register()}
                    errors={errors.name}
                    className={""}
                  />
                </div>
                <div className="my-6 space-y-2 flex justify-center">
                  <button
                    type="submit"
                    style={{ background: "#1163E6" }}
                    className="button button__md button__primary w-full"
                  >
                    Continue
                  </button>
                </div>
              </form>

              <div className="mt-6 space-y-2 flex justify-center">
                <a
                  style={{ background: "#fff", color: "#000" }}
                  className="button button__md w-full"
                  href="/sign-in"
                >
                  Back to login
                </a>
              </div>
            </div>
          </div>
        </section>
      </main>
    </>
  );
}

export default ForgotPasswordPage;
