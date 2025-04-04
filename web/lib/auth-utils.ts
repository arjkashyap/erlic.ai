import { AuthResponse } from "@/types/auth";
import { RequestCookie } from "next/dist/compiled/@edge-runtime/cookies";

export async function checkAuth(sessionCookie: RequestCookie | undefined): Promise<AuthResponse> {
  if (!sessionCookie) {
    return { authenticated: false, user: null };
  }

  try {
    const response = await fetch("http://localhost:8080/api/auth/me", {
      headers: {
        Cookie: `${sessionCookie.name}=${sessionCookie.value}`,
      },
    });

    if (response.ok) {
      return await response.json();
    }

    return { authenticated: false, user: null };
  } catch (error) {
    console.error("Error checking auth:", error);
    return { authenticated: false, user: null };
  }
} 