import { StyleSheet, View } from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { useTheme } from "../theme/ThemeContext";
import BotonGrandote from "./BotonGrandote";

export default function SettingsButtons() {
    const { colors } = useTheme();

    return (
        <View style={styles.container}>
            <View style={styles.botones}>
                <BotonGrandote to="/AddMoney" color={colors.primary}>
                    <Ionicons name="person-outline" size={120} color={colors.primary} />
                </BotonGrandote>
                <BotonGrandote to="/AddMoney" color={colors.primary}>
                    <Ionicons name="git-branch-outline" size={120} color={colors.primary} />
                </BotonGrandote>
            </View>
            <View style={styles.botones}>
                <BotonGrandote to="/(tabs)/(settings)/AddCategory" color={colors.primary}>
                    <Ionicons name="layers-outline" size={120} color={colors.primary} />
                </BotonGrandote>
                <BotonGrandote to="/AddMoney" color={colors.primary}>
                    <Ionicons name="list-outline" size={120} color={colors.primary} />
                </BotonGrandote>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: "center",
        backgroundColor: "#121212",
        alignItems: "center",
    },

    botones: {
        marginBottom: 20,
        flexDirection: "row",
        gap: 20,
    },

});

