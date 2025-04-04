import Link from "next/link"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Shield } from "lucide-react"
import { FcGoogle } from "react-icons/fc"
import { SiGithub } from "react-icons/si"
import { FaMicrosoft } from "react-icons/fa"

export default function LoginPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <Card className="w-full max-w-md">
        <CardHeader className="space-y-1">
          <div className="flex justify-center mb-4">
            <Link href="/" className="flex items-center space-x-2">
              <Shield className="h-6 w-6 text-purple-700" />
              <span className="text-xl font-bold">Erlic</span>
            </Link>
          </div>
          <CardTitle className="text-2xl text-center">Welcome back</CardTitle>
          <CardDescription className="text-center">Enter your credentials to access your account</CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="space-y-2">
            <label htmlFor="email" className="text-sm font-medium">
              Email
            </label>
            <Input id="email" type="email" placeholder="name@company.com" />
          </div>
          <div className="space-y-2">
            <div className="flex items-center justify-between">
              <label htmlFor="password" className="text-sm font-medium">
                Password
              </label>
              <Link href="/forgot-password" className="text-sm text-purple-700 hover:underline">
                Forgot password?
              </Link>
            </div>
            <Input id="password" type="password" />
          </div>
        </CardContent>
        <CardFooter className="flex flex-col space-y-4">
          <Button className="w-full bg-purple-700 hover:bg-purple-800">Sign In</Button>
          
          <div className="relative my-4">
            <div className="absolute inset-0 flex items-center">
              <span className="w-full border-t" />
            </div>
            <div className="relative flex justify-center text-xs uppercase">
              <span className="bg-background px-2 text-muted-foreground">
                Or continue with
              </span>
            </div>
          </div>

          <div className="grid grid-cols-3 gap-2">
            <Button variant="outline" className="w-full" asChild>
              <Link href="http://localhost:8080/api/auth/google">
                <FcGoogle className="mr-2 h-4 w-4" />
                <span className="sr-only">Google</span>
              </Link>
            </Button>
            <Button variant="outline" className="w-full" >
              <FaMicrosoft className="mr-2 h-4 w-4 text-blue-600" />
              <span className="sr-only">Microsoft</span>
            </Button>
            <Button variant="outline" className="w-full" >
              <SiGithub className="mr-2 h-4 w-4" />
              <span className="sr-only">GitHub</span>
            </Button>
          </div>

          <div className="text-center text-sm">
            Don't have an account?{" "}
            <Link href="/signup" className="text-purple-700 hover:underline">
              Sign up
            </Link>
          </div>
        </CardFooter>
      </Card>
    </div>
  )
}

