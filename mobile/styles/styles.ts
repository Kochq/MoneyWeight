import { StyleSheet } from "react-native";

export const AddMoneyStyles = (colors: any) => {
    return StyleSheet.create({
        container: {
            flex: 1,
            padding: 20,
            backgroundColor: colors.background,
        },

        card: {
            backgroundColor: colors.surface,
            minWidth: "100%",
            borderRadius: 15,
            padding: 20,
            boxShadow: "0 0 10px rgba(0, 0, 0, 0.1)",
        },

        inputContainer: {
            marginBottom: 16,
        },

        label: {
            fontSize: 14,
            color: colors.textSecondary,
            marginBottom: 8,
            fontWeight: "500",
        },

        input: {
            backgroundColor: colors.background,
            borderRadius: 10,
            padding: 12,
            fontSize: 16,
            color: colors.textPrimary,
            borderWidth: 1,
            borderColor: colors.border,
        },

        picker: {
            backgroundColor: colors.background,
            borderRadius: 10,
            marginBottom: 16,
            borderWidth: 1,
            borderColor: colors.border,
        },

        dateContainer: {
            flexDirection: "row",
            alignItems: "center",
            marginBottom: 20,
            backgroundColor: colors.background,
            borderRadius: 10,
            padding: 12,
            borderWidth: 1,
            borderColor: colors.border,
        },

        dateText: {
            flex: 1,
            color: colors.textPrimary,
            fontSize: 16,
        },

        dateButton: {
            padding: 8,
            borderRadius: 8,
            marginLeft: 8,
            backgroundColor: colors.primary + "20",
        },

        addButton: {
            backgroundColor: colors.primary,
            borderRadius: 12,
            padding: 16,
            alignItems: "center",
            marginTop: 10,
        },

        addButtonText: {
            color: colors.surface,
            fontSize: 16,
            fontWeight: "600",
        },

        moneyInput: {
            fontSize: 24,
            textAlign: "center",
            color: colors.primary,
            fontWeight: "bold",
        },

        titleInput: {
            textAlign: "center",
            color: colors.textPrimary,
        },
    });
};

export const TransactionStyles = (colors: any) => {
    return StyleSheet.create({
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
            backgroundColor: colors.surface,
        },

        title: {
            maxWidth: "60%",
            minWidth: "30%",
            color: colors.textPrimary,
            fontSize: 16,
        },

        desc: {
            flexDirection: "column",
            alignItems: "flex-end",
        },

        money: {
            color: colors.moneyIn,
            fontSize: 16,
        },

        end: {
            flexDirection: "row",
            alignItems: "center",
            gap: 10,
        },
    });
};

export const TotalCardStyles = (colors: any) => {
    return StyleSheet.create({
        container: {
            borderWidth: 0.1,
            width: "100%",
            height: 100,
            marginBottom: 20,
            padding: 10,
            boxShadow: "0 0 10px rgba(0, 0, 0, 0.1)",
            flexDirection: "row",
            borderRadius: 6,
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
            fontSize: 20,
            color: colors.textPrimary,
        },

        money: {
            fontSize: 16,
        },
    });
};
