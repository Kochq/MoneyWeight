import { useState } from "react";
import { Text, View, StyleSheet, Switch } from "react-native";

export const TotalCard = () => {
    const [isEnabled, setIsEnabled] = useState(false);
    const toggleSwitch = () => setIsEnabled((previousState) => !previousState);
    return (
        <View style={styles.container}>
            <View style={styles.status}>
                <Text style={styles.title}>Total money</Text>
                <Text style={styles.money}>$ 1000</Text>
            </View>
            <Switch
                style={styles.switch}
                trackColor={{ false: "#767577", true: "#81b0ff" }}
                thumbColor={isEnabled ? "#f5dd4b" : "#f4f3f4"}
                ios_backgroundColor="#3e3e3e"
                onValueChange={toggleSwitch}
                value={isEnabled}
            />
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        borderWidth: 0.1,
        width: "100%",
        height: 100,
        marginBottom: 20,
        padding: 10,
        flexDirection: "row",
        borderRadius: 6,
        borderColor: "#fff",
    },

    switch: {
        height: "50%",
        width: "50%",
    },

    status: {
        height: "100%",
        width: "50%",
        flexDirection: "column",
        justifyContent: "space-between",
    },

    title: {
        color: "#fff",
        fontSize: 20,
    },

    money: {
        color: "#0F0",
        fontSize: 16,
    },
});
