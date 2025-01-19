import { Transaction } from "@/types";
import { View, Text, Pressable } from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { useTheme } from "../theme/ThemeContext";
import { Link } from "expo-router";
import { TransactionStyles } from "@/styles/styles";

export default function TransactionCard({ t }: { t: Transaction }) {
    const { colors } = useTheme();
    const day = t.date.split(" ")[0];
    const styles = TransactionStyles(colors);

    return (
        <Link
            asChild
            href={{
                pathname: "/transaction/[id]",
                params: { id: t.id },
            }}
        >
            <Pressable>
                <View style={styles.container}>
                    <View style={styles.titles}>
                        <Ionicons
                            name="ellipse"
                            size={24}
                            color={colors.textPrimary}
                        />
                        <Text
                            style={styles.title}
                            numberOfLines={1}
                            ellipsizeMode="tail"
                        >
                            {t.title}
                        </Text>
                    </View>
                    <View style={styles.end}>
                        <View style={styles.desc}>
                            <Text style={styles.money}>{`$ ${t.amount}`}</Text>
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
