import { Stack } from "expo-router";
import React from "react";

export default function HomeLayout() {
    return (
        <Stack
            screenOptions={{
                headerShown: true,
                headerStyle: {
                    backgroundColor: "#f4511e",
                },
                headerTintColor: "#fff",
                headerTitleStyle: {
                    fontWeight: "bold",
                },
            }}
        >
            <Stack.Screen name="index" />
            <Stack.Screen name="details" />
        </Stack>
    );
}
