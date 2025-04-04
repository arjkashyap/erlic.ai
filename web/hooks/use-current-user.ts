import { useEffect, useState } from "react";
import { User } from "@/types/auth"; // Assuming this type definition exists

export function useCurrentUser() {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  console.log("useCurrentUser hook initialized");

  useEffect(() => {
    console.log("useEffect in useCurrentUser triggered");
    
    const fetchUser = async () => {
      console.log("Starting fetchUser");
      try {
        console.log("Making API request to /api/auth/me...");
        const response = await fetch("/api/auth/me", {
          credentials: "include",
        });

        console.log("Response status:", response.status);
        
        if (response.ok) {
          const data = await response.json();
          console.log("Raw response data:", data);
          
          if (data.authenticated && data.user) {
            console.log("Setting user data:", data.user);
            setUser(data.user);
          } else {
            console.log("No user data in response");
          }
        } else {
          console.log("Response not OK:", response.status);
          const errorText = await response.text();
          console.log("Error response:", errorText);
        }
      } catch (error) {
        console.error("Error fetching user:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchUser();
  }, []);

  return { user, loading };
}
