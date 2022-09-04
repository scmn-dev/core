import React from "react";
import Head from "next/head";
import * as yup from "yup";
import Api from "api";
import CryptoJS from "crypto-js";
import { useRouter } from "next/router";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { ToastContainer, toast } from "react-toastify";
import { TextInput } from "@components/text-input";
import { ErrorMsg } from "components/errors";
import { LogoSection } from "components/logo";

const schema = yup.object().shape({
  password: yup.string().required(),
  passwordConfirm: yup
    .string()
    .oneOf([yup.ref("password"), null], "Passwords must match")
    .required(),
});

function FinalStepPage() {
  const router = useRouter();
  const { register, handleSubmit, errors } = useForm({
    resolver: yupResolver(schema),
  });

  const registerAPI = ({ name, email, password }) => {
    return Api.post(`/auth/signup`, {
      name,
      email,
      master_password: CryptoJS.SHA256(password).toString(),
    })
      .then((data) => Promise.resolve(data))
      .catch((err) => {
        toast(<ErrorMsg msg={err.response.data.message} />);

        return Promise.reject(err);
      });
  };

  const onSubmit = async ({ password }) => {
    let name = localStorage.getItem("name");
    let email = localStorage.getItem("email");

    registerAPI({ name, email, password })
      .then(() => {
        router.push("/confirmed");
      })
      .catch((err) => console.error(err));
  };

  return (
    <>
      <ToastContainer hideProgressBar />

      <main className="flex flex-col md:flex-row-reverse md:h-screen">
        <Head>
          <title>Final Step - $PASSWORD_MANAGER_NAME</title>
          <link rel="icon" href="$LOGO_URL" />
        </Head>
        <LogoSection />

        <section className="justify-center px-4 md:px-0 md:flex md:w-2/3">
          <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
            <h2 className="text-lg font-medium md:text-2xl">
              Set master password
            </h2>
            <p className="text-sm pt-2 text-gray-400">
              You can't change your master password later and we do not store
              your master password anywhere. Please keep it secret and do not
              forget.
            </p>
            <div className="my-4">
              <form onSubmit={handleSubmit(onSubmit)} method="POST">
                <div className="mb-3">
                  <TextInput
                    className=""
                    placeholder=""
                    label="Master Password"
                    name="password"
                    type="password"
                    register={register()}
                    errors={errors.password}
                  />
                </div>

                <div className="mb-3">
                  <TextInput
                    className=""
                    placeholder=""
                    label="Master Password Verify"
                    name="passwordConfirm"
                    type="password"
                    register={register()}
                    errors={errors.password}
                  />
                </div>

                <div className="mt-6 space-y-2 flex justify-center">
                  <button
                    type="submit"
                    style={{ background: "#1163E6" }}
                    className="button button__md button__primary w-full"
                  >
                    Create My Account
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

export default FinalStepPage;
