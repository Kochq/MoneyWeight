import React, { useEffect, useState } from "react";
import {
    View,
    Text,
    TouchableOpacity,
    Modal,
    Pressable,
    StyleSheet,
} from "react-native";
import { useTheme } from "../theme/ThemeContext";
import { Ionicons } from "@expo/vector-icons";
import { Accounts } from "@/types";

const AccountSelector = () => {
    const { colors } = useTheme();
    const [modalVisible, setModalVisible] = useState(false);
    const [selectedAccount, setSelectedAccount] = useState("Todas");
    const [accounts, setAccounts] = useState<Accounts[]>([]);
    const apiUrl = process.env.EXPO_PUBLIC_API_BASE;

    useEffect(() => {
        const getAccounts = async () => {
            try {
                const url = apiUrl + "/api/accounts";
                const response = await fetch(url);
                const data = await response.json();
                setAccounts(data.data);
            } catch (e) {
                console.error("Error al obtener categorías:", e);
            }
        };

        getAccounts();
    }, [apiUrl]);

    const styles = StyleSheet.create({
        container: {
            paddingHorizontal: 8,
        },
        selectorButton: {
            flexDirection: "row",
            alignItems: "center",
            gap: 4,
        },
        selectorText: {
            color: colors.textPrimary,
            fontSize: 18,
        },
        modalOverlay: {
            flex: 1,
            backgroundColor: "rgba(0, 0, 0, 0.5)",
        },
        modalContent: {
            marginTop: "auto",
            backgroundColor: colors.surface,
            borderTopLeftRadius: 24,
            borderTopRightRadius: 24,
        },
        modalHeader: {
            padding: 16,
        },
        modalIndicator: {
            width: 48,
            height: 4,
            backgroundColor: colors.border,
            borderRadius: 9999,
            alignSelf: "center",
            marginBottom: 16,
        },
        modalTitle: {
            color: colors.textPrimary,
            fontSize: 20,
            fontWeight: "bold",
            marginBottom: 16,
        },
        accountItem: {
            paddingVertical: 16,
            paddingHorizontal: 8,
            borderBottomWidth: 1,
            borderBottomColor: colors.border,
        },
        accountText: {
            color: colors.textPrimary,
            fontSize: 18,
        },
        icon: {
            marginLeft: 4,
        },
    });

    return (
        <View style={styles.container}>
            <TouchableOpacity
                onPress={() => setModalVisible(true)}
                style={styles.selectorButton}
            >
                <Text style={styles.selectorText}>{selectedAccount}</Text>
                <Ionicons
                    name="chevron-down"
                    size={20}
                    color={colors.textPrimary}
                    style={styles.icon}
                />
            </TouchableOpacity>

            <Modal
                animationType="slide"
                transparent={true}
                visible={modalVisible}
                onRequestClose={() => setModalVisible(false)}
            >
                <Pressable
                    style={styles.modalOverlay}
                    onPress={() => setModalVisible(false)}
                >
                    <View style={styles.modalContent}>
                        <View style={styles.modalHeader}>
                            <View style={styles.modalIndicator} />
                            <Text style={styles.modalTitle}>
                                Seleccionar Cuenta
                            </Text>

                            {accounts.map((account) => (
                                <TouchableOpacity
                                    key={account.id}
                                    style={styles.accountItem}
                                    onPress={() => {
                                        setSelectedAccount(account.name);
                                        setModalVisible(false);
                                    }}
                                >
                                    <Text
                                        style={[
                                            styles.accountText,
                                            {
                                                color:
                                                    selectedAccount ===
                                                    account.name
                                                        ? colors.primary
                                                        : colors.textPrimary,
                                            },
                                        ]}
                                    >
                                        {account.name}
                                    </Text>
                                </TouchableOpacity>
                            ))}
                        </View>
                    </View>
                </Pressable>
            </Modal>
        </View>
    );
};

export default AccountSelector;
