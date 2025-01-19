import { useState } from "react";
import { Text, View, Switch } from "react-native";
import { useTheme } from "../theme/ThemeContext";
import { TotalCardStyles } from "@/styles/styles";

export default function TotalCard() {
    const { colors, toggleTheme } = useTheme();
    const [isEnabled, setIsEnabled] = useState(false);
    const styles = TotalCardStyles(colors);
    const toggleSwitch = () => {
        setIsEnabled((previousState) => !previousState);
        toggleTheme();
    };

    return (
        <View style={[styles.container, { backgroundColor: colors.surface }]}>
            <View style={styles.status}>
                <Text style={styles.title}>Total money</Text>
                <Text style={styles.title}>$ 1000</Text>
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
}
