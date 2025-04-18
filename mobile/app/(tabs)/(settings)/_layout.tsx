// (tabs)/(settings)/_layout.tsx
import { Stack } from "expo-router";
import React from "react";

export default function SettingsLayout() {
    return (
        <Stack
            screenOptions={{
                headerShown: true,
                headerStyle: {
                    backgroundColor: "#121212",
                },
                headerTintColor: "#fff",
                headerTitleStyle: {
                    fontWeight: "bold",
                },
            }}
        >
            <Stack.Screen
                name="index"
                options={{
                    title: "Settings",
                }}
            />
            <Stack.Screen
                name="AddCategory"
                options={{
                    title: "Add Category",
                }}
            />
        </Stack>
    );
}
