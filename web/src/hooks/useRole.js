import { useRouter } from "next/router";
import { useEffect } from "react";
export const useRole = (role) => {
  const router = useRouter();
  useEffect(() => {
    if (role === 0) {
      router.push("/login");
    }
    if (role === 1) {
      router.push("/admin");
    }
    if (role > 5) {
      router.push("/mesero");
    }
  }, [role]);
};
