import { Stack } from "expo-router";
import React from "react";
import { StyleSheet } from "react-native";

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
            <Stack.Screen name="details" />
            <Stack.Screen name="AddMoney" />
        </Stack>
    );
}
