import { Stack } from "expo-router";
import { StatusBar } from "expo-status-bar";
import { SafeAreaProvider } from "react-native-safe-area-context";
import { ThemeProvider } from "../theme/ThemeContext";

export default function RootLayout() {
    return (
        <SafeAreaProvider>
            <ThemeProvider>
                <StatusBar style="auto" />
                <Stack
                    screenOptions={{
                        headerShown: false,
                    }}
                >
                    <Stack.Screen name="(tabs)" />
                </Stack>
            </ThemeProvider>
        </SafeAreaProvider>
    );
}
