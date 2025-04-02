import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Shield, Cloud, Database, Server, PlusCircle, CheckCircle, RefreshCw, Settings } from "lucide-react"

export default function ConnectorsPage() {
  return (
    <div className="flex-1 p-8 space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-3xl font-bold tracking-tight">Connectors</h1>
        <Button>
          <PlusCircle className="mr-2 h-4 w-4" />
          Add Connector
        </Button>
      </div>

      <Tabs defaultValue="active" className="space-y-4">
        <TabsList>
          <TabsTrigger value="active">Active Connectors</TabsTrigger>
          <TabsTrigger value="available">Available Connectors</TabsTrigger>
          <TabsTrigger value="settings">Settings</TabsTrigger>
        </TabsList>

        <TabsContent value="active" className="space-y-4">
          <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
            {/* Active Directory Connector */}
            <Card>
              <CardHeader className="pb-2">
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-2">
                    <Shield className="h-5 w-5 text-purple-700" />
                    <CardTitle className="text-lg">Active Directory</CardTitle>
                  </div>
                  <div className="flex items-center space-x-1 bg-green-100 text-green-700 px-2 py-1 rounded-full text-xs">
                    <CheckCircle className="h-3 w-3" />
                    <span>Active</span>
                  </div>
                </div>
                <CardDescription>Windows Active Directory</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="text-sm space-y-2">
                  <div className="flex justify-between">
                    <span className="text-muted-foreground">Server:</span>
                    <span className="font-medium">ad.company.local</span>
                  </div>
                  <div className="flex justify-between">
                    <span className="text-muted-foreground">Domain:</span>
                    <span className="font-medium">COMPANY</span>
                  </div>
                  <div className="flex justify-between">
                    <span className="text-muted-foreground">Last sync:</span>
                    <span className="font-medium">10 minutes ago</span>
                  </div>
                </div>
              </CardContent>
              <CardFooter className="flex justify-between pt-2">
                <Button variant="outline" size="sm">
                  <RefreshCw className="mr-2 h-3 w-3" />
                  Sync
                </Button>
                <Button variant="outline" size="sm">
                  <Settings className="mr-2 h-3 w-3" />
                  Configure
                </Button>
              </CardFooter>
            </Card>
          </div>
        </TabsContent>

        <TabsContent value="available" className="space-y-4">
          <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
            {/* Azure AD Connector */}
            <Card>
              <CardHeader className="pb-2">
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-2">
                    <Cloud className="h-5 w-5 text-blue-600" />
                    <CardTitle className="text-lg">Azure AD</CardTitle>
                  </div>
                </div>
                <CardDescription>Microsoft Azure Active Directory</CardDescription>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-muted-foreground mb-4">
                  Connect to Azure AD to manage users, groups, and applications.
                </p>
              </CardContent>
              <CardFooter>
                <Button className="w-full">Configure</Button>
              </CardFooter>
            </Card>

            {/* Google Workspace Connector */}
            <Card>
              <CardHeader className="pb-2">
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-2">
                    <Cloud className="h-5 w-5 text-yellow-600" />
                    <CardTitle className="text-lg">Google Workspace</CardTitle>
                  </div>
                </div>
                <CardDescription>Google Workspace (G Suite)</CardDescription>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-muted-foreground mb-4">
                  Connect to Google Workspace to manage users and services.
                </p>
              </CardContent>
              <CardFooter>
                <Button className="w-full">Configure</Button>
              </CardFooter>
            </Card>

            {/* AWS Connector */}
            <Card>
              <CardHeader className="pb-2">
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-2">
                    <Cloud className="h-5 w-5 text-orange-600" />
                    <CardTitle className="text-lg">AWS</CardTitle>
                  </div>
                </div>
                <CardDescription>Amazon Web Services</CardDescription>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-muted-foreground mb-4">Connect to AWS to manage resources and services.</p>
              </CardContent>
              <CardFooter>
                <Button className="w-full">Configure</Button>
              </CardFooter>
            </Card>

            {/* LDAP Connector */}
            <Card>
              <CardHeader className="pb-2">
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-2">
                    <Database className="h-5 w-5 text-gray-600" />
                    <CardTitle className="text-lg">LDAP</CardTitle>
                  </div>
                </div>
                <CardDescription>Lightweight Directory Access Protocol</CardDescription>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-muted-foreground mb-4">
                  Connect to LDAP servers to manage directory information.
                </p>
              </CardContent>
              <CardFooter>
                <Button className="w-full">Configure</Button>
              </CardFooter>
            </Card>

            {/* SCCM Connector */}
            <Card>
              <CardHeader className="pb-2">
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-2">
                    <Server className="h-5 w-5 text-purple-600" />
                    <CardTitle className="text-lg">SCCM</CardTitle>
                  </div>
                </div>
                <CardDescription>System Center Configuration Manager</CardDescription>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-muted-foreground mb-4">
                  Connect to SCCM to manage devices and applications.
                </p>
              </CardContent>
              <CardFooter>
                <Button className="w-full">Configure</Button>
              </CardFooter>
            </Card>
          </div>
        </TabsContent>

        <TabsContent value="settings" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Connector Settings</CardTitle>
              <CardDescription>Configure global settings for all connectors</CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="sync-interval">Sync Interval (minutes)</Label>
                <Input id="sync-interval" type="number" defaultValue="30" />
              </div>

              <div className="space-y-2">
                <Label htmlFor="timeout">Connection Timeout (seconds)</Label>
                <Input id="timeout" type="number" defaultValue="60" />
              </div>

              <div className="space-y-2">
                <Label htmlFor="retry-attempts">Retry Attempts</Label>
                <Input id="retry-attempts" type="number" defaultValue="3" />
              </div>
            </CardContent>
            <CardFooter>
              <Button>Save Settings</Button>
            </CardFooter>
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  )
}

