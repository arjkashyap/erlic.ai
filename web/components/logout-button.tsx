import { LogOut } from 'lucide-react';
// import { useAuth } from '@/hooks/use-auth';
import { Button } from '@/components/ui/button';

interface LogoutButtonProps {
  variant?: 'default' | 'destructive' | 'outline' | 'secondary' | 'ghost' | 'link';
  size?: 'default' | 'sm' | 'lg' | 'icon';
}

export function LogoutButton({ variant = 'ghost', size = 'default' }: LogoutButtonProps) {
  // const { logout } = useAuth();

  return (
    <Button variant={variant} size={size} >
      <LogOut className="mr-2 h-4 w-4" />
      <span>Log out</span>
    </Button>
  );
}