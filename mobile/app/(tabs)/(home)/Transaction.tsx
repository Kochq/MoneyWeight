import { Transaction } from "@/types";
import { View, Text, StyleSheet } from "react-native";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { useTheme } from "../../../theme/ThemeContext";

export default function TransactionCard({ t }: { t: Transaction }) {
    const { colors } = useTheme();
    const day = t.date.split(" ")[0];

    return (
        <View style={[styles.container, { backgroundColor: colors.surface }]}>
            <View style={styles.titles}>
                <MaterialIcons
                    name="category"
                    size={24}
                    color={colors.textPrimary}
                />
                <View style={styles.title}>
                    <Text
                        style={{ color: colors.textPrimary }}
                        numberOfLines={1}
                        ellipsizeMode="tail"
                    >
                        {t.title}
                    </Text>
                </View>
            </View>
            <View style={styles.desc}>
                <Text style={{ color: colors.moneyIn }}>{`$ ${t.amount}`}</Text>
                <Text style={{ color: colors.textPrimary }}>{day}</Text>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    titles: {
        flexDirection: "row",
        alignItems: "center",
        gap: 10,
    },

    container: {
        borderRadius: 5,
        flexDirection: "row",
        boxShadow: "0 0 5px rgba(0, 0, 0, 0.1)",
        borderWidth: 0.1,
        width: "100%",
        paddingVertical: 10,
        paddingHorizontal: 20,
        alignItems: "center",
        justifyContent: "space-between",
        marginBottom: 10,
    },

    title: {
        width: "60%",
    },

    desc: {
        flexDirection: "column",
        alignItems: "flex-end",
    },
});
