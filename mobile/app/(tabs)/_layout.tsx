import { Tabs } from "expo-router";
import { useTheme } from "../../theme/ThemeContext";
import { Ionicons } from "@expo/vector-icons";

export default function TabLayout() {
    const { colors } = useTheme();
    return (
        <Tabs
            screenOptions={{
                tabBarActiveTintColor: colors.primary,
                headerShown: false,
            }}
        >
            <Tabs.Screen
                name="(home)"
                options={{
                    tabBarIcon: ({ color }) => (
                        <Ionicons name="home-sharp" size={24} color={color} />
                    ),
                    tabBarActiveBackgroundColor: colors.background,
                    tabBarInactiveBackgroundColor: colors.background,
                    title: "Home",
                }}
            />
            <Tabs.Screen
                name="(settings)"
                options={{
                    tabBarIcon: ({ color }) => (
                        <Ionicons name="settings-sharp" size={24} color={color} />
                    ),
                    tabBarActiveBackgroundColor: colors.background,
                    tabBarInactiveBackgroundColor: colors.background,
                    title: "Settings",
                }}
            />
        </Tabs>
    );
}
