// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.CostManagement.V20190401Preview
{
    /// <summary>
    /// Show costs accumulated over time.
    /// </summary>
    [EnumType]
    public readonly struct AccumulatedType : IEquatable<AccumulatedType>
    {
        private readonly string _value;

        private AccumulatedType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AccumulatedType @True { get; } = new AccumulatedType("true");
        public static AccumulatedType @False { get; } = new AccumulatedType("false");

        public static bool operator ==(AccumulatedType left, AccumulatedType right) => left.Equals(right);
        public static bool operator !=(AccumulatedType left, AccumulatedType right) => !left.Equals(right);

        public static explicit operator string(AccumulatedType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AccumulatedType other && Equals(other);
        public bool Equals(AccumulatedType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The category of the budget, whether the budget tracks cost or usage.
    /// </summary>
    [EnumType]
    public readonly struct CategoryType : IEquatable<CategoryType>
    {
        private readonly string _value;

        private CategoryType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CategoryType Cost { get; } = new CategoryType("Cost");
        public static CategoryType Usage { get; } = new CategoryType("Usage");

        public static bool operator ==(CategoryType left, CategoryType right) => left.Equals(right);
        public static bool operator !=(CategoryType left, CategoryType right) => !left.Equals(right);

        public static explicit operator string(CategoryType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CategoryType other && Equals(other);
        public bool Equals(CategoryType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Chart type of the main view in Cost Analysis. Required.
    /// </summary>
    [EnumType]
    public readonly struct ChartType : IEquatable<ChartType>
    {
        private readonly string _value;

        private ChartType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ChartType Area { get; } = new ChartType("Area");
        public static ChartType Line { get; } = new ChartType("Line");
        public static ChartType StackedColumn { get; } = new ChartType("StackedColumn");
        public static ChartType GroupedColumn { get; } = new ChartType("GroupedColumn");
        public static ChartType Table { get; } = new ChartType("Table");

        public static bool operator ==(ChartType left, ChartType right) => left.Equals(right);
        public static bool operator !=(ChartType left, ChartType right) => !left.Equals(right);

        public static explicit operator string(ChartType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ChartType other && Equals(other);
        public bool Equals(ChartType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The name of the aggregation function to use.
    /// </summary>
    [EnumType]
    public readonly struct FunctionType : IEquatable<FunctionType>
    {
        private readonly string _value;

        private FunctionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FunctionType Sum { get; } = new FunctionType("Sum");

        public static bool operator ==(FunctionType left, FunctionType right) => left.Equals(right);
        public static bool operator !=(FunctionType left, FunctionType right) => !left.Equals(right);

        public static explicit operator string(FunctionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FunctionType other && Equals(other);
        public bool Equals(FunctionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The granularity of rows in the report.
    /// </summary>
    [EnumType]
    public readonly struct GranularityType : IEquatable<GranularityType>
    {
        private readonly string _value;

        private GranularityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static GranularityType Daily { get; } = new GranularityType("Daily");
        public static GranularityType Monthly { get; } = new GranularityType("Monthly");

        public static bool operator ==(GranularityType left, GranularityType right) => left.Equals(right);
        public static bool operator !=(GranularityType left, GranularityType right) => !left.Equals(right);

        public static explicit operator string(GranularityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GranularityType other && Equals(other);
        public bool Equals(GranularityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// KPI type (Forecast, Budget).
    /// </summary>
    [EnumType]
    public readonly struct KpiTypeType : IEquatable<KpiTypeType>
    {
        private readonly string _value;

        private KpiTypeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KpiTypeType Forecast { get; } = new KpiTypeType("Forecast");
        public static KpiTypeType Budget { get; } = new KpiTypeType("Budget");

        public static bool operator ==(KpiTypeType left, KpiTypeType right) => left.Equals(right);
        public static bool operator !=(KpiTypeType left, KpiTypeType right) => !left.Equals(right);

        public static explicit operator string(KpiTypeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KpiTypeType other && Equals(other);
        public bool Equals(KpiTypeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Metric to use when displaying costs.
    /// </summary>
    [EnumType]
    public readonly struct MetricType : IEquatable<MetricType>
    {
        private readonly string _value;

        private MetricType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MetricType ActualCost { get; } = new MetricType("ActualCost");
        public static MetricType AmortizedCost { get; } = new MetricType("AmortizedCost");
        public static MetricType AHUB { get; } = new MetricType("AHUB");

        public static bool operator ==(MetricType left, MetricType right) => left.Equals(right);
        public static bool operator !=(MetricType left, MetricType right) => !left.Equals(right);

        public static explicit operator string(MetricType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MetricType other && Equals(other);
        public bool Equals(MetricType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The comparison operator.
    /// </summary>
    [EnumType]
    public readonly struct NotificationOperatorType : IEquatable<NotificationOperatorType>
    {
        private readonly string _value;

        private NotificationOperatorType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static NotificationOperatorType EqualTo { get; } = new NotificationOperatorType("EqualTo");
        public static NotificationOperatorType GreaterThan { get; } = new NotificationOperatorType("GreaterThan");
        public static NotificationOperatorType GreaterThanOrEqualTo { get; } = new NotificationOperatorType("GreaterThanOrEqualTo");

        public static bool operator ==(NotificationOperatorType left, NotificationOperatorType right) => left.Equals(right);
        public static bool operator !=(NotificationOperatorType left, NotificationOperatorType right) => !left.Equals(right);

        public static explicit operator string(NotificationOperatorType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NotificationOperatorType other && Equals(other);
        public bool Equals(NotificationOperatorType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The operator to use for comparison.
    /// </summary>
    [EnumType]
    public readonly struct OperatorType : IEquatable<OperatorType>
    {
        private readonly string _value;

        private OperatorType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OperatorType In { get; } = new OperatorType("In");
        public static OperatorType Contains { get; } = new OperatorType("Contains");

        public static bool operator ==(OperatorType left, OperatorType right) => left.Equals(right);
        public static bool operator !=(OperatorType left, OperatorType right) => !left.Equals(right);

        public static explicit operator string(OperatorType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OperatorType other && Equals(other);
        public bool Equals(OperatorType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Data type to show in view.
    /// </summary>
    [EnumType]
    public readonly struct PivotTypeType : IEquatable<PivotTypeType>
    {
        private readonly string _value;

        private PivotTypeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PivotTypeType Dimension { get; } = new PivotTypeType("Dimension");
        public static PivotTypeType TagKey { get; } = new PivotTypeType("TagKey");

        public static bool operator ==(PivotTypeType left, PivotTypeType right) => left.Equals(right);
        public static bool operator !=(PivotTypeType left, PivotTypeType right) => !left.Equals(right);

        public static explicit operator string(PivotTypeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PivotTypeType other && Equals(other);
        public bool Equals(PivotTypeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Has type of the column to group.
    /// </summary>
    [EnumType]
    public readonly struct ReportConfigColumnType : IEquatable<ReportConfigColumnType>
    {
        private readonly string _value;

        private ReportConfigColumnType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ReportConfigColumnType Tag { get; } = new ReportConfigColumnType("Tag");
        public static ReportConfigColumnType Dimension { get; } = new ReportConfigColumnType("Dimension");

        public static bool operator ==(ReportConfigColumnType left, ReportConfigColumnType right) => left.Equals(right);
        public static bool operator !=(ReportConfigColumnType left, ReportConfigColumnType right) => !left.Equals(right);

        public static explicit operator string(ReportConfigColumnType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ReportConfigColumnType other && Equals(other);
        public bool Equals(ReportConfigColumnType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the report. Usage represents actual usage, forecast represents forecasted data and UsageAndForecast represents both usage and forecasted data. Actual usage and forecasted data can be differentiated based on dates.
    /// </summary>
    [EnumType]
    public readonly struct ReportType : IEquatable<ReportType>
    {
        private readonly string _value;

        private ReportType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ReportType Usage { get; } = new ReportType("Usage");

        public static bool operator ==(ReportType left, ReportType right) => left.Equals(right);
        public static bool operator !=(ReportType left, ReportType right) => !left.Equals(right);

        public static explicit operator string(ReportType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ReportType other && Equals(other);
        public bool Equals(ReportType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
    /// </summary>
    [EnumType]
    public readonly struct TimeGrainType : IEquatable<TimeGrainType>
    {
        private readonly string _value;

        private TimeGrainType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TimeGrainType Monthly { get; } = new TimeGrainType("Monthly");
        public static TimeGrainType Quarterly { get; } = new TimeGrainType("Quarterly");
        public static TimeGrainType Annually { get; } = new TimeGrainType("Annually");

        public static bool operator ==(TimeGrainType left, TimeGrainType right) => left.Equals(right);
        public static bool operator !=(TimeGrainType left, TimeGrainType right) => !left.Equals(right);

        public static explicit operator string(TimeGrainType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TimeGrainType other && Equals(other);
        public bool Equals(TimeGrainType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
    /// </summary>
    [EnumType]
    public readonly struct TimeframeType : IEquatable<TimeframeType>
    {
        private readonly string _value;

        private TimeframeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TimeframeType WeekToDate { get; } = new TimeframeType("WeekToDate");
        public static TimeframeType MonthToDate { get; } = new TimeframeType("MonthToDate");
        public static TimeframeType YearToDate { get; } = new TimeframeType("YearToDate");
        public static TimeframeType Custom { get; } = new TimeframeType("Custom");

        public static bool operator ==(TimeframeType left, TimeframeType right) => left.Equals(right);
        public static bool operator !=(TimeframeType left, TimeframeType right) => !left.Equals(right);

        public static explicit operator string(TimeframeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TimeframeType other && Equals(other);
        public bool Equals(TimeframeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
