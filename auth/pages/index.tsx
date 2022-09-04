import { useRouter } from "next/router";
import React, { useEffect } from "react";

export default function HomePage() {
  const router = useRouter();

  useEffect(() => {
    setTimeout(() => {
      router.push("/sign-up");
    }, 1);
  }, []);

  return <div></div>;
}
