import React, { useContext } from "react";
import Head from "next/head";
import { useForm } from "react-hook-form";
import { TextInput } from "@components/text-input";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import AppContext from "store/form-type";
import Api from "api";
import { ToastContainer, toast } from "react-toastify";
import { useRouter } from "next/router";
import { ErrorMsg } from "components/errors";
import { LogoSection } from "components/logo";

const schema = yup.object().shape({
  name: yup
    .string()
    .matches(/^.[a-zA-ZıİçÇşŞğĞÜüÖö ]+$/, {
      message: "Your name can only contain alphabetic characters",
    })
    .required(),
  email: yup.string().email().required(),
});

export default function Page() {
  const router = useRouter();

  const { register, handleSubmit, errors } = useForm({
    resolver: yupResolver(schema),
  });

  const store = useContext(AppContext);
  store.changeFormType("free");

  const createCodeAPI = async ({ name, email }) => {
    return await Api.post(`/auth/code`, {
      name,
      email,
    })
      .then((data) => Promise.resolve(data))
      .catch((err) => {
        toast(
          <ErrorMsg
            msg="This email is already taken. Please use another one."
            messages={[err.response.data.message]}
          />
        );

        return Promise.reject(err);
      });
  };

  const onSubmit = async ({ name, email }) => {
    localStorage.setItem("name", name);
    localStorage.setItem("email", email);

    createCodeAPI({ name, email })
      .then(() => {
        router.push("/code");
      })
      .catch((err) => console.error(err));
  };

  return (
    <>
      <ToastContainer hideProgressBar />

      <main className="flex flex-col md:flex-row-reverse md:h-screen">
        <Head>
          <title>Sign Up - $PASSWORD_MANAGER_NAME</title>
          <link rel="icon" href="$LOGO_URL" />
        </Head>
        <LogoSection />

        <section className="justify-center px-4 md:px-0 md:flex md:w-2/3">
          <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
            <h2 className="text-lg font-medium md:text-2xl">
              Create an Account
            </h2>
            {/* <p className="text-sm pt-2 text-gray-400">
              Already have an account?{" "}
              <a className="font-medium" href="/sign-in">
                Sign in
              </a>
            </p> */}
            <div className="my-4">
              <form onSubmit={handleSubmit(onSubmit)} method="POST">
                <div className="mb-3">
                  <label
                    htmlFor="username"
                    className="block text-sm font-medium text-neutral-400"
                  >
                    Username
                  </label>
                  <TextInput
                    name="name"
                    placeholder="Username"
                    register={register()}
                    errors={errors.name}
                    className={""}
                  />
                </div>

                <div className="mb-3">
                  <label
                    htmlFor="email"
                    className="block text-sm font-medium text-neutral-400"
                  >
                    Email address
                  </label>
                  <TextInput
                    name="email"
                    placeholder="Email"
                    type="email"
                    register={register()}
                    errors={errors.email}
                    className={""}
                  />
                </div>

                <div className="flex justify-between mb-4">
                  <p className="text-small transition-colors">
                    By registering, you agree to the processing of your personal
                    data by <strong>$PASSWORD_MANAGER_NAME</strong> as described in the{" "}
                    <a
                      href="$PASSWORD_MANAGER_PRIVACY_POLICY_URL"
                      target="_blank"
                      rel="noopener noreferrer"
                    >
                      Privacy Policy
                    </a>
                    .
                  </p>
                </div>

                <div className="mt-6 space-y-2 flex justify-center">
                  <button
                    type="submit"
                    style={{ background: "#1163E6" }}
                    className="button button__md button__primary w-full"
                  >
                    Sign up
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
