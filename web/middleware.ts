import { NextRequest, NextResponse } from "next/server";
import { cookies } from "next/headers"; // Keep for other potential uses, but not direct decoding
import { checkAuth } from "./lib/auth-utils";

const protectedRoutes = ["/dashboard"];
const publicRoutes = ["/login", "/signup", "/"];

const AUTH_CHECK_URL = "http://localhost:8080/api/auth/me";

export default async function middleware(req: NextRequest) {
  const path = req.nextUrl.pathname;
  const isProtectedRoute = protectedRoutes.some((route) =>
    path.startsWith(route)
  );
  const isPublicRoute = publicRoutes.includes(path);

  // Get the session cookie to pass it along
  // Note: Fetch usually forwards cookies automatically for same-origin,
  // but explicitly getting it might be needed depending on setup or for debugging.
  // For cross-origin (localhost:3000 -> localhost:8080), CORS headers
  // (Access-Control-Allow-Credentials: true, Access-Control-Allow-Origin: http://localhost:3000)
  // MUST be set correctly on your Go backend.
  const sessionCookie = (await cookies()).get("_gothic_session");

  const authResult = await checkAuth(sessionCookie);

  // 4. Redirect to /login if the user is not authenticated and accessing a protected route
  if (isProtectedRoute && !authResult.authenticated) {
    console.log(
      `Middleware: Redirecting unauthenticated user from protected route ${path} to /login`
    );
    return NextResponse.redirect(new URL("/login", req.url));
  }

  // 5. Redirect to /dashboard if the user is authenticated and trying to access a public-only route like /login
  if (isPublicRoute && authResult.authenticated) {
    console.log(
      `Middleware: Redirecting authenticated user from public route ${path} to /dashboard`
    );
    return NextResponse.redirect(new URL("/dashboard", req.url));
  }

  // 6. Allow the request to proceed
  return NextResponse.next();
}

// Routes Middleware should not run on
export const config = {
  matcher: ["/((?!api|_next/static|_next/image|favicon.ico|.*\\.png$).*)"], // Added favicon.ico
};
