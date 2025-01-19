import { Stack } from "expo-router";
import React from "react";

export default function HomeLayout() {
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
            <Stack.Screen name="index" />
            <Stack.Screen name="details/[id]" />
            <Stack.Screen name="transaction/[id]" />
            <Stack.Screen name="AddMoney" />
        </Stack>
    );
}
