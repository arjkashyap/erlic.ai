
import Link from "next/link"
import { Button } from "@/components/ui/button"
import { ArrowRight, CheckCircle, Shield, Zap, MessageSquare } from "lucide-react"


export default function LandingPage() {


  
  return (
    <div className="flex flex-col min-h-screen">
      {/* Hero Section */}
      <header className="bg-gradient-to-r from-purple-700 to-indigo-800 text-white">
        <div className="container mx-auto px-4 py-6">
          <nav className="flex justify-between items-center">
            <div className="flex items-center space-x-2">
              <Shield className="h-8 w-8" />
              <span className="text-2xl font-bold">Erlic</span>
            </div>
            <div className="hidden md:flex space-x-6">
              <Link href="#features" className="hover:text-purple-200 transition">
                Features
              </Link>
              <Link href="#how-it-works" className="hover:text-purple-200 transition">
                How it works
              </Link>
              <Link href="#pricing" className="hover:text-purple-200 transition">
                Pricing
              </Link>
            </div>
            <div className="flex space-x-4">
              <Link href="/login">
                <Button variant="outline" className="bg-white/10 text-white border-white hover:bg-white hover:text-purple-700">
                  Login
                </Button>
              </Link>
              <Link href="/signup">
              <Button variant="outline" className="bg-white/10 text-white border-white hover:bg-white hover:text-purple-700">
                  Sign Up
                </Button>
              </Link>
            </div>
          </nav>

          <div className="py-20 text-center">
            <h1 className="text-4xl md:text-6xl font-bold mb-6">AI-Powered IT Help Desk</h1>
            <p className="text-xl md:text-2xl mb-10 max-w-3xl mx-auto">
              Automate your IT administration tasks with AI. Manage tickets, reset passwords, and configure systems with
              simple chat commands.
            </p>
            <div className="flex flex-col sm:flex-row justify-center gap-4">
              <Link href="/signup">
                <Button size="lg" className="bg-white text-purple-700 hover:bg-purple-100 transition-all duration-200 hover:scale-105 hover:shadow-lg">
                  Get Started <ArrowRight className="ml-2 h-5 w-5" />
                </Button>
              </Link>
              <Link href="#demo">
                <Button size="lg" className="bg-white text-purple-700 hover:bg-purple-100 transition-all duration-200 hover:scale-105 hover:shadow-lg">
                  See Demo <ArrowRight className="ml-2 h-5 w-5" />
                </Button>
              </Link>
            </div>
          </div>
        </div>
      </header>

      {/* Features Section */}
      <section id="features" className="py-20 bg-gray-50">
        <div className="container mx-auto px-4">
          <h2 className="text-3xl md:text-4xl font-bold text-center mb-16">Powerful Features for IT Administrators</h2>

          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-10">
            <div className="bg-white p-8 rounded-xl shadow-md">
              <div className="bg-purple-100 p-3 rounded-full w-fit mb-6">
                <MessageSquare className="h-6 w-6 text-purple-700" />
              </div>
              <h3 className="text-xl font-bold mb-4">AI-Powered Assistant</h3>
              <p className="text-gray-600">
                Execute IT tasks through natural language commands. Reset passwords, manage accounts, and more with
                simple chat prompts.
              </p>
            </div>

            <div className="bg-white p-8 rounded-xl shadow-md">
              <div className="bg-purple-100 p-3 rounded-full w-fit mb-6">
                <Zap className="h-6 w-6 text-purple-700" />
              </div>
              <h3 className="text-xl font-bold mb-4">Smart Connectors</h3>
              <p className="text-gray-600">
                Seamlessly integrate with Active Directory, Azure, GCP, and other services through our intelligent
                connector system.
              </p>
            </div>

            <div className="bg-white p-8 rounded-xl shadow-md">
              <div className="bg-purple-100 p-3 rounded-full w-fit mb-6">
                <CheckCircle className="h-6 w-6 text-purple-700" />
              </div>
              <h3 className="text-xl font-bold mb-4">Ticket Management</h3>
              <p className="text-gray-600">
                Track and manage support tickets efficiently. Assign, prioritize, and resolve issues all from one
                dashboard.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* How It Works */}
      <section id="how-it-works" className="py-20">
        <div className="container mx-auto px-4">
          <h2 className="text-3xl md:text-4xl font-bold text-center mb-16">How Erlic Works</h2>

          <div className="grid md:grid-cols-3 gap-8">
            <div className="text-center">
              <div className="bg-purple-700 text-white rounded-full w-12 h-12 flex items-center justify-center mx-auto mb-6">
                1
              </div>
              <h3 className="text-xl font-bold mb-4">Configure Connectors</h3>
              <p className="text-gray-600">
                Set up connections to your IT systems like Active Directory, Azure, or GCP.
              </p>
            </div>

            <div className="text-center">
              <div className="bg-purple-700 text-white rounded-full w-12 h-12 flex items-center justify-center mx-auto mb-6">
                2
              </div>
              <h3 className="text-xl font-bold mb-4">Manage Tickets</h3>
              <p className="text-gray-600">Receive and track support requests from your organization's users.</p>
            </div>

            <div className="text-center">
              <div className="bg-purple-700 text-white rounded-full w-12 h-12 flex items-center justify-center mx-auto mb-6">
                3
              </div>
              <h3 className="text-xl font-bold mb-4">Use AI Commands</h3>
              <p className="text-gray-600">Resolve issues with simple chat commands that automate routine tasks.</p>
            </div>
          </div>
        </div>
      </section>

      {/* Pricing Section */}
      <section id="pricing" className="py-20 bg-gray-50">
        <div className="container mx-auto px-4">
          <h2 className="text-3xl md:text-4xl font-bold text-center mb-16">Simple, Transparent Pricing</h2>

          <div className="grid md:grid-cols-3 gap-8 max-w-5xl mx-auto">
            <div className="bg-white p-8 rounded-xl shadow-md border border-gray-200">
              <h3 className="text-xl font-bold mb-2">Starter</h3>
              <p className="text-gray-600 mb-6">For small IT teams</p>
              <p className="text-4xl font-bold mb-6">
                $29<span className="text-lg text-gray-600">/month</span>
              </p>

              <ul className="space-y-3 mb-8">
                <li className="flex items-center">
                  <CheckCircle className="h-5 w-5 text-green-500 mr-2" />
                  <span>Up to 3 admin users</span>
                </li>
                <li className="flex items-center">
                  <CheckCircle className="h-5 w-5 text-green-500 mr-2" />
                  <span>Basic connectors</span>
                </li>
                <li className="flex items-center">
                  <CheckCircle className="h-5 w-5 text-green-500 mr-2" />
                  <span>100 tickets/month</span>
                </li>
              </ul>

              <Button className="w-full">Get Started</Button>
            </div>

            <div className="bg-purple-700 text-white p-8 rounded-xl shadow-md transform scale-105">
              <div className="absolute top-0 right-0 bg-yellow-400 text-xs font-bold px-3 py-1 rounded-bl-lg rounded-tr-lg text-purple-900">
                POPULAR
              </div>
              <h3 className="text-xl font-bold mb-2">Professional</h3>
              <p className="text-purple-200 mb-6">For growing businesses</p>
              <p className="text-4xl font-bold mb-6">
                $79<span className="text-lg text-purple-200">/month</span>
              </p>

              <ul className="space-y-3 mb-8">
                <li className="flex items-center">
                  <CheckCircle className="h-5 w-5 text-green-300 mr-2" />
                  <span>Up to 10 admin users</span>
                </li>
                <li className="flex items-center">
                  <CheckCircle className="h-5 w-5 text-green-300 mr-2" />
                  <span>All connectors</span>
                </li>
                <li className="flex items-center">
                  <CheckCircle className="h-5 w-5 text-green-300 mr-2" />
                  <span>500 tickets/month</span>
                </li>
                <li className="flex items-center">
                  <CheckCircle className="h-5 w-5 text-green-300 mr-2" />
                  <span>Advanced AI capabilities</span>
                </li>
              </ul>

              <Button className="w-full bg-white text-purple-700 hover:bg-purple-100">Get Started</Button>
            </div>

            <div className="bg-white p-8 rounded-xl shadow-md border border-gray-200">
              <h3 className="text-xl font-bold mb-2">Enterprise</h3>
              <p className="text-gray-600 mb-6">For large organizations</p>
              <p className="text-4xl font-bold mb-6">
                $199<span className="text-lg text-gray-600">/month</span>
              </p>

              <ul className="space-y-3 mb-8">
                <li className="flex items-center">
                  <CheckCircle className="h-5 w-5 text-green-500 mr-2" />
                  <span>Unlimited admin users</span>
                </li>
                <li className="flex items-center">
                  <CheckCircle className="h-5 w-5 text-green-500 mr-2" />
                  <span>Custom connectors</span>
                </li>
                <li className="flex items-center">
                  <CheckCircle className="h-5 w-5 text-green-500 mr-2" />
                  <span>Unlimited tickets</span>
                </li>
                <li className="flex items-center">
                  <CheckCircle className="h-5 w-5 text-green-500 mr-2" />
                  <span>Priority support</span>
                </li>
              </ul>

              <Button className="w-full">Contact Sales</Button>
            </div>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-900 text-white py-12">
        <div className="container mx-auto px-4">
          <div className="grid md:grid-cols-4 gap-8">
            <div>
              <div className="flex items-center space-x-2 mb-6">
                <Shield className="h-6 w-6" />
                <span className="text-xl font-bold">Erlic</span>
              </div>
              <p className="text-gray-400">AI-powered IT help desk solution for modern businesses.</p>
            </div>

            <div>
              <h4 className="text-lg font-bold mb-4">Product</h4>
              <ul className="space-y-2">
                <li>
                  <Link href="#features" className="text-gray-400 hover:text-white">
                    Features
                  </Link>
                </li>
                <li>
                  <Link href="#pricing" className="text-gray-400 hover:text-white">
                    Pricing
                  </Link>
                </li>
                <li>
                  <Link href="#" className="text-gray-400 hover:text-white">
                    Integrations
                  </Link>
                </li>
              </ul>
            </div>

            <div>
              <h4 className="text-lg font-bold mb-4">Resources</h4>
              <ul className="space-y-2">
                <li>
                  <Link href="#" className="text-gray-400 hover:text-white">
                    Documentation
                  </Link>
                </li>
                <li>
                  <Link href="#" className="text-gray-400 hover:text-white">
                    Blog
                  </Link>
                </li>
                <li>
                  <Link href="#" className="text-gray-400 hover:text-white">
                    Support
                  </Link>
                </li>
              </ul>
            </div>

            <div>
              <h4 className="text-lg font-bold mb-4">Company</h4>
              <ul className="space-y-2">
                <li>
                  <Link href="#" className="text-gray-400 hover:text-white">
                    About
                  </Link>
                </li>
                <li>
                  <Link href="#" className="text-gray-400 hover:text-white">
                    Careers
                  </Link>
                </li>
                <li>
                  <Link href="#" className="text-gray-400 hover:text-white">
                    Contact
                  </Link>
                </li>
              </ul>
            </div>
          </div>

          <div className="border-t border-gray-800 mt-12 pt-8 text-center text-gray-400">
            <p>&copy; {new Date().getFullYear()} Erlic. All rights reserved.</p>
          </div>
        </div>
      </footer>
    </div>
  )
}

