import { Transaction } from "@/types";
import { View, Text, StyleSheet, Pressable } from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { useTheme } from "../theme/ThemeContext";
import { Link } from "expo-router";

export default function TransactionCard({ t }: { t: Transaction }) {
    const { colors } = useTheme();
    const day = t.date.split(" ")[0];

    return (
        <Link
            asChild
            href={{
                pathname: "/transaction/[id]",
                params: { id: t.id },
            }}
        >
            <Pressable>
                <View
                    style={[
                        styles.container,
                        { backgroundColor: colors.surface },
                    ]}
                >
                    <View style={styles.titles}>
                        <Ionicons
                            name="ellipse"
                            size={24}
                            color={colors.textPrimary}
                        />
                        <View style={styles.title}>
                            <Text
                                style={{
                                    color: colors.textPrimary,
                                    fontSize: 16,
                                }}
                                numberOfLines={1}
                                ellipsizeMode="tail"
                            >
                                {t.title}
                            </Text>
                        </View>
                    </View>
                    <View style={styles.end}>
                        <View style={styles.desc}>
                            <Text
                                style={{ color: colors.moneyIn, fontSize: 16 }}
                            >{`$ ${t.amount}`}</Text>
                            <Text
                                style={{
                                    color: colors.textPrimary,
                                    fontSize: 12,
                                }}
                            >
                                {day}
                            </Text>
                        </View>
                        <Ionicons
                            name="chevron-forward"
                            size={24}
                            color={colors.textPrimary}
                        />
                    </View>
                </View>
            </Pressable>
        </Link>
    );
}

const styles = StyleSheet.create({
    titles: {
        flexDirection: "row",
        alignItems: "center",
        gap: 10,
    },

    container: {
        borderRadius: 6,
        flexDirection: "row",
        boxShadow: "0 0 10px rgba(0, 0, 0, 0.1)",
        borderWidth: 0.1,
        minWidth: "100%",
        maxWidth: "100%",
        paddingVertical: 12,
        paddingHorizontal: 12,
        alignItems: "center",
        justifyContent: "space-between",
        marginBottom: 10,
    },

    title: {
        maxWidth: "60%",
        minWidth: "30%",
    },

    desc: {
        flexDirection: "column",
        alignItems: "flex-end",
    },

    end: {
        flexDirection: "row",
        alignItems: "center",
        gap: 10,
    },
});
